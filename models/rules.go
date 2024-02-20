package models

import (
	"github.com/juandunbar/immunity/database"
)

type RulesStore struct {
	DB database.Database
}

type Rule struct {
	ID          int    `db:"id"`
	Query       string `db:"query"`
	Description string `db:"description"`
	Action      string `db:"action"`
	LastUsed    string `db:"last_used"`
	Disabled    bool   `db:"disabled"`
}

func NewRulesStore(db database.Database) *RulesStore {
	return &RulesStore{DB: db}
}

func (r *RulesStore) GetRuleList() ([]Rule, error) {
	rules := make([]Rule, 0)
	err := r.DB.Query(&rules, "SELECT * FROM rules")
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (r *RulesStore) GetRuleListMap() (map[int]Rule, error) {
	rules, err := r.GetRuleList()
	if err != nil {
		return nil, err
	}
	rulesMap := make(map[int]Rule, 0)
	for _, v := range rules {
		rulesMap[v.ID] = v
	}
	return rulesMap, nil
}
