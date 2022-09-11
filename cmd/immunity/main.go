package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	// import our custom processor
	"github.com/juandunbar/immunity/api"
	"github.com/juandunbar/immunity/benthos"
	_ "github.com/juandunbar/immunity/processors/rules"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// handle syscalls to stop application
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
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
		if err := benthos.RunStream(ctx); err != nil {
			log.Println(err)
		}
	}()
	// block while we wait for any interrupt or term signals
	<-done
	cancel()
	// wait for processes to close
	wg.Wait()
	log.Println("goodbye, immunity has shutdown")
	os.Exit(0)
}
