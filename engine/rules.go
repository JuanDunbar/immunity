package engine

import (
	"errors"
	"fmt"
	"time"

	"github.com/timbray/quamina"
)

type Rule struct {
	ID          string
	Query       string
	Description string
	Action      string
	LastUsed    time.Time
}

type RulesEngine struct {
	rulesList map[string]Rule
	matcher   *quamina.Quamina
}

// TODO get rules from database
var list = map[string]Rule{
	"sadf23425": {
		ID:          "sadf2342",
		Query:       "",
		Description: "test rule one",
		Action:      "fishy",
		LastUsed:    time.Now(),
	},
	"ert324234": {
		ID:          "ert324234",
		Query:       "",
		Description: "test rule two",
		Action:      "spam",
		LastUsed:    time.Now(),
	},
	"zxcv42389": {
		ID:          "zxcv42389",
		Query:       "",
		Description: "test rule three",
		Action:      "fishy",
		LastUsed:    time.Now(),
	},
}

func NewRulesEngine() (*RulesEngine, error) {
	matcher, err := quamina.New()
	if err != nil {
		return nil, err
	}
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
	return r.matcher.MatchesForEvent(event)
}

func (r *RulesEngine) GetRule(id string) (*Rule, error) {
	rule, ok := list[id]
	if ok != true {
		return nil, errors.New(fmt.Sprintf("failed to get data for rule: %s", id))
	}
	return &rule, nil
}
