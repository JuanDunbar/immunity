package models

import (
	"github.com/juandunbar/immunity/config"
	"github.com/juandunbar/immunity/database"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TestType struct {
	One int `db:"one_column"`
	Two int `db:"two_column"`
}

func TestRulesModal(t *testing.T) {
	testConfig, err := config.LoadConfig()
	assert.Nil(t, err)
	db := database.NewDatabase()
	err = db.Connect(testConfig)
	assert.Nil(t, err)
	//store := NewRulesStore(db)
	err = db.Execute(`INSERT INTO rules (query, description, action, last_used, disabled) VALUES (:query,:description,:action,:last_used,:disabled)`,
		map[string]interface{}{
			"query":       `{"name":"testname"}`,
			"description": "test rule",
			"action":      "alert",
			"last_used":   time.Now(),
			"disabled":    false,
		})
	assert.Nil(t, err)
	testQuery := make([]Rule, 0)
	err = db.Query(&testQuery, "SELECT * FROM public.rules")
	assert.Nil(t, err)
}
