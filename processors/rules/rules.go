package rules

import (
	"context"
	"github.com/juandunbar/immunity/events"
	"time"

	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/goccy/go-json"

	"github.com/juandunbar/immunity/engine"
)

func init() {
	configSpec := service.NewConfigSpec()

	constructor := func(config *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {
		return newRulesProcessor(mgr.Logger()), nil
	}

	err := service.RegisterProcessor("rules", configSpec, constructor)
	if err != nil {
		// TODO fail gracefully
		panic(err)
	}
}

type rulesProcessor struct {
	logger      *service.Logger
	rulesEngine *engine.RulesEngine
}

func newRulesProcessor(logger *service.Logger) *rulesProcessor {
	rulesEngine, _ := engine.NewRulesEngine()
	return &rulesProcessor{
		logger:      logger,
		rulesEngine: rulesEngine,
	}
}

func (r *rulesProcessor) Process(ctx context.Context, m *service.Message) (service.MessageBatch, error) {
	event, err := m.AsBytes()
	if err != nil {
		return nil, err
	}
	// add our original event to the output, so we can store in elasticsearch
	outputMessages := []*service.Message{m}
	// run our event through our rules engine to find any matching rules
	matches, err := r.rulesEngine.Match(event)
	for _, match := range matches {
		rule, err := r.rulesEngine.GetRule(match.(string))
		if err != nil {
			return nil, err
		}
		// create new output event that we can filter on for actions
		suspiciousActivity := events.SuspiciousActivity{
			Event:     "suspicious_activity",
			Data:      string(event),
			Rule:      rule.ID,
			Action:    rule.Action,
			Timestamp: time.Now(),
		}
		newMessage, _ := json.Marshal(suspiciousActivity)
		outputMessages = append(outputMessages, service.NewMessage(newMessage))
	}

	return outputMessages, nil
}

func (r *rulesProcessor) Close(ctx context.Context) error {
	return nil
}
