package chain

import log "github.com/sirupsen/logrus"

var (
	processChains = make(map[string]*chain)
)

type Chain interface {
	Process() error
}

type Handler func(event any) (bool, error)
type Constructor func(data []byte) (any, error)

type chain struct {
	event       any
	constructor Constructor
	handlers    []Handler
}

func (c *chain) Process() error {
	for _, handler := range c.handlers {
		result, err := handler(c.event)
		if err != nil {
			return err
		}
		if result == true {
			break
		}
	}
	return nil
}

func LoadChain(chainName string, constructor Constructor, handlers ...Handler) {
	pChain := chain{
		constructor: constructor,
		handlers:    handlers,
	}
	processChains[chainName] = &pChain
}

func GetChain(chainName string, eventData []byte) (Chain, bool) {
	pChain, ok := processChains[chainName]
	if ok == false {
		return nil, false
	}
	event, err := pChain.constructor(eventData)
	if err != nil {
		log.WithField("chainName", chainName).
			WithError(err).
			Error("type constructor failed for event")
		return nil, false
	}
	pChain.event = event

	return pChain, true
}
