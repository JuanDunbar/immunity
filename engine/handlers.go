package engine

import (
	log "github.com/sirupsen/logrus"

	"github.com/juandunbar/immunity/chain"
	// side import all chain types to load them into our chain map
	_ "github.com/juandunbar/immunity/events/account_created"
)

type HandlersEngine struct{}

func NewHandlersEngine() (*HandlersEngine, error) {
	return &HandlersEngine{}, nil
}

func (h *HandlersEngine) Process(eventType string, eventData []byte) {
	handlersChain, ok := chain.GetChain(eventType, eventData)
	if ok != true {
		log.WithFields(log.Fields{
			"@service":   "immunity",
			"event_type": eventType,
		}).Warningf("failed to find handlers for event")
		return
	}
	// run through all our process chains handler funcs
	err := handlersChain.Process()
	if err != nil {
		log.WithFields(log.Fields{
			"@service":   "immunity",
			"event_type": eventType,
		}).WithError(err).Error("error in process chain")
	}
}
