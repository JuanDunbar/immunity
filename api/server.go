package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("This is an Immunity test endpoint"))
}

func Run(ctx context.Context) {
	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/rules", TestEndpoint).Methods("GET")

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
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	// Block until we receive our signal.
	<-ctx.Done()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
