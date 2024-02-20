package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"

	"github.com/juandunbar/immunity/engine"
	"github.com/juandunbar/immunity/models"
)

type rulesController struct {
	Store *models.RulesStore
}

func (rc *rulesController) GetRules(w http.ResponseWriter, r *http.Request) {
	rules, err := rc.Store.GetRuleList()
	resp, err := json.Marshal(rules)
	if err != nil {
		http.Error(w, "failed to retrieve rules list", 500)
		return
	}
	w.Write(resp)
}

func (rc *rulesController) PostRules(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte(fmt.Sprintf("not implemented")))
}

type rulesStreamController struct {
	RulesEngine *engine.RulesEngine
}

func (sc *rulesStreamController) ProcessRulesStream(w http.ResponseWriter, r *http.Request) {
	var event []byte
	var err error

	if r.Body == nil {
		http.Error(w, "request body required", 400)
		return
	}

	event, err = io.ReadAll(r.Body)
	if err != nil {
		log.WithField("@service", "immunity").
			WithError(err).
			Error("error reading request body")
		http.Error(w, "internal server error", 500)
		return
	}
	defer r.Body.Close()

	eventType := r.Header.Get("event_type")

	log.WithFields(log.Fields{
		"event_type": eventType,
		"event_data": string(event),
	}).Info("event received")

	matches, err := sc.RulesEngine.Match(event)
	for _, match := range matches {
		rule, err := sc.RulesEngine.GetRule(match.(int))
		if err != nil {
			log.WithField("@service", "immunity").
				WithError(err).
				Error("error getting rule from match")
			http.Error(w, "internal server error", 500)
			return
		}
		// TODO call handler based on the rules defined action
		log.WithFields(log.Fields{
			"rule_id":     rule.ID,
			"rule_query":  rule.Query,
			"rule_action": rule.Action,
			"rule_desc":   rule.Description,
			"event_type":  eventType,
			"event_data":  string(event),
		}).Warn("rule match")
	}
}
