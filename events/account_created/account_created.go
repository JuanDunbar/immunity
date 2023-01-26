package account_created

import (
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"

	"github.com/juandunbar/immunity/chain"
)

const EventName = "account_created"

func init() {
	chain.LoadChain(EventName, NewAccountCreated, StepOne, StepTwo)
}

type AccountCreated struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IpAddress string `json:"ip_address"`
}

func NewAccountCreated(data []byte) (any, error) {
	accountCreated := new(AccountCreated)
	err := json.Unmarshal(data, accountCreated)
	if err != nil {
		return nil, err
	}
	return accountCreated, nil
}

func StepOne(event any) (bool, error) {
	accountCreated := event.(*AccountCreated)
	log.WithFields(log.Fields{
		"handler":    "StepOne",
		"first_name": accountCreated.FirstName,
		"last_name":  accountCreated.LastName,
		"ip_address": accountCreated.IpAddress,
	}).Info("new account_created event processed")
	return false, nil
}

func StepTwo(event any) (bool, error) {
	accountCreated := event.(*AccountCreated)
	log.WithFields(log.Fields{
		"handler":    "StepTwo",
		"first_name": accountCreated.FirstName,
		"last_name":  accountCreated.LastName,
		"ip_address": accountCreated.IpAddress,
	}).Info("new account_created event processed")
	return true, nil
}
