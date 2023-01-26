package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := mux.NewRouter()
	// Add your routes as needed
	rc := new(rulesController)
	r.HandleFunc("/rules", rc.GetRules).Methods("GET")
	r.HandleFunc("/rules/add", rc.PostRules).Methods("GET")

	srv := &http.Server{
		Addr: ":8181",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).WithFields(log.Fields{
				"@service": "immunity",
			}).Error("failed to run api server")
		}
	}()
	// Block until we receive our signal.
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.WithError(err).WithFields(log.Fields{
			"@service": "immunity",
		}).Error("failed to shutdown api server")
	}
}
