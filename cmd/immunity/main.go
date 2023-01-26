package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/juandunbar/immunity/api"
	"github.com/juandunbar/immunity/benthos"
	"github.com/juandunbar/immunity/database"
	_ "github.com/juandunbar/immunity/processors/handlers"
	_ "github.com/juandunbar/immunity/processors/rules"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// handle syscalls to stop application
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	// init our database connection
	if err := database.Connect(ctx); err != nil {
		log.WithFields(log.Fields{
			"@service": "immunity",
		}).WithError(err).Error("error connecting to mongo database")
	}
	// create workgroup to wait for processes to close
	wg := &sync.WaitGroup{}
	// start our rules api server
	wg.Add(1)
	go func() {
		defer wg.Done()
		api.Run(ctx)
	}()
	// start our event stream
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := benthos.RunStream(ctx); err != nil && err != context.Canceled {
			log.WithFields(log.Fields{
				"@service": "immunity",
			}).WithError(err).Error("error running rules event stream")
			done <- syscall.SIGTERM
		}
	}()
	// block while we wait for any interrupt or term signals
	<-done
	cancel()
	if err := database.Disconnect(); err != nil {
		log.WithFields(log.Fields{
			"@service": "immunity",
		}).WithError(err).Error("error disconnecting from mongo database")
	}
	// wait for processes to close
	wg.Wait()
	log.WithFields(log.Fields{
		"@service": "immunity",
	}).Info("goodbye, immunity has shutdown")
}
