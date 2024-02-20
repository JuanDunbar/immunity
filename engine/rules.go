package engine

import (
	"errors"
	"fmt"

	"github.com/timbray/quamina"

	"github.com/juandunbar/immunity/models"
)

type RulesEngine struct {
	rulesList map[int]models.Rule
	matcher   *quamina.Quamina
}

func NewRulesEngine(store *models.RulesStore) (*RulesEngine, error) {
	matcher, err := quamina.New()
	if err != nil {
		return nil, err
	}
	// get our rules from the DB as a lookup table
	list, err := store.GetRuleListMap()
	// load rules into our matcher
	for k, v := range list {
		err = matcher.AddPattern(k, v.Query)
		if err != nil {
			return nil, err
		}
	}
	return &RulesEngine{
		rulesList: list,
		matcher:   matcher,
	}, nil
}

func (r *RulesEngine) Match(event []byte) ([]quamina.X, error) {
	// TODO this is not safe for concurrent use
	return r.matcher.MatchesForEvent(event)
}

func (r *RulesEngine) GetRule(id int) (*models.Rule, error) {
	rule, ok := r.rulesList[id]
	if ok != true {
		return nil, errors.New(fmt.Sprintf("failed to get data for rule: %s", id))
	}
	return &rule, nil
}
