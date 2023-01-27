package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/juandunbar/immunity/config"
)

type Api interface {
	Run(errCh chan error)
	Shutdown() error
}

type api struct {
	server *http.Server
}

func NewApiServer(c *config.Config) Api {
	r := mux.NewRouter()
	// Add your routes as needed
	rc := new(rulesController)
	r.HandleFunc("/rules", rc.GetRules).Methods("GET")
	r.HandleFunc("/rules/add", rc.PostRules).Methods("GET")

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", c.Api.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	return &api{
		server: srv,
	}
}

func (api *api) Run(errCh chan error) {
	if err := api.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		errCh <- err
	}
}

func (api *api) Shutdown() error {
	return api.server.Shutdown(context.TODO())
}
