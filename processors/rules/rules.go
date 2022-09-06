package rules

import (
	"context"
	"github.com/juandunbar/immunity/engine"

	"github.com/benthosdev/benthos/v4/public/service"
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
	matches, err := r.rulesEngine.Match(event)
	for _, match := range matches {
		rule, err := r.rulesEngine.GetRule(match.(string))
		if err != nil {
			return nil, err
		}
		// TODO run the rule action
		r.logger.Debugf("event matched rule", rule, string(event))
	}

	return []*service.Message{m}, nil
}

func (r *rulesProcessor) Close(ctx context.Context) error {
	return nil
}
