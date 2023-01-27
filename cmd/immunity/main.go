package main

import (
	log "github.com/sirupsen/logrus"

	immunity "github.com/juandunbar/immunity/service"
)

func main() {
	svc := immunity.NewService()
	if err := svc.Start(); err != nil {
		log.WithField("@service", "immunity").
			WithError(err).
			Fatal("Failed to start service!")
	}
}
