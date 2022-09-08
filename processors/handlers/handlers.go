package handlers

import (
	"context"
	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/juandunbar/immunity/engine"
)

func init() {
	configSpec := service.NewConfigSpec()

	constructor := func(config *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {
		return newHandlersProcessor(mgr.Logger()), nil
	}

	err := service.RegisterProcessor("handlers", configSpec, constructor)
	if err != nil {
		// TODO fail gracefully
		panic(err)
	}
}

type handlersProcessor struct {
	logger         *service.Logger
	handlersEngine *engine.HandlersEngine
}

func newHandlersProcessor(logger *service.Logger) *handlersProcessor {
	handlersEngine, _ := engine.NewHandlersEngine()
	return &handlersProcessor{
		logger:         logger,
		handlersEngine: handlersEngine,
	}
}

func (r *handlersProcessor) Process(ctx context.Context, m *service.Message) (service.MessageBatch, error) {
	// need to get our event type from event headers
	//eventType, _ := m.MetaGet("event_type")

	// add our original event to the output, so we can store in elasticsearch
	outputMessages := []*service.Message{m}
	// run our event through our rules engine to find any matching rules

	return outputMessages, nil
}

func (r *handlersProcessor) Close(ctx context.Context) error {
	return nil
}
