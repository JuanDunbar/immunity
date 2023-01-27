package service

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/juandunbar/immunity/api"
	"github.com/juandunbar/immunity/benthos"
	"github.com/juandunbar/immunity/config"
	"github.com/juandunbar/immunity/database"
	_ "github.com/juandunbar/immunity/processors/handlers"
	_ "github.com/juandunbar/immunity/processors/rules"
)

type Service interface {
	Start() error
	Stop()
}

type service struct {
	Cfg    *config.Config
	DB     database.Database
	Api    api.Api
	ErrCh  chan error
	SigCh  chan os.Signal
	StopWG sync.WaitGroup
	mutex  sync.Mutex
}

func NewService() Service {
	return &service{}
}

func (s *service) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error

	s.mutex.Lock()
	defer s.mutex.Unlock()
	// create channels
	s.ErrCh = make(chan error, 1)
	s.SigCh = make(chan os.Signal, 10)
	// load our config
	s.Cfg, err = config.LoadConfig()
	if err != nil {
		return errors.Wrapf(err, "while loading config")
	}
	// connect our database
	s.DB = database.NewDatabase()
	err = s.DB.Connect(s.Cfg)
	if err != nil {
		return errors.Wrapf(err, "while connecting to database")
	}
	// signal notifications and handler
	signal.Notify(s.SigCh, syscall.SIGINT, syscall.SIGTERM)
	s.StopWG.Add(1)
	go s.handleSignals()
	// start our rules api server
	s.Api = api.NewApiServer(s.Cfg)
	s.StopWG.Add(1)
	go s.startApi()
	// start our event stream
	s.StopWG.Add(1)
	go s.startEventStream(ctx)

	// block and wait for any errors
	err = <-s.ErrCh
	cancel()
	s.Stop()
	s.StopWG.Wait()

	if err != nil {
		log.WithField("@service", "immunity").
			WithError(err).
			Fatal("Error while shutting down!")
	}
	log.WithField("@service", "immunity").Info("immunity shutdown!")

	return nil
}

func (s *service) Stop() {
	var err error
	if err = s.DB.Disconnect(); err != nil {
		log.WithField("@service", "immunity").
			WithError(err).
			Error("Error shutting down database!")
	}
	if err = s.Api.Shutdown(); err != nil {
		log.WithField("@service", "immunity").
			WithError(err).
			Error("Error shutting down HTTP server!")
	}
}

func (s *service) handleSignals() {
	defer s.StopWG.Done()
	select {
	case sig := <-s.SigCh:
		switch sig {
		case syscall.SIGTERM:
			fallthrough
		case syscall.SIGINT:
			s.Stop()
			return
		}
	}
}

func (s *service) startApi() {
	defer s.StopWG.Done()
	s.Api.Run(s.ErrCh)
}

func (s *service) startEventStream(ctx context.Context) {
	defer s.StopWG.Done()
	if err := benthos.RunStream(ctx); err != nil && err != context.Canceled {
		s.ErrCh <- err
	}
}
