package chain

import (
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestType struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewTestType(data []byte) (any, error) {
	testType := new(TestType)
	err := json.Unmarshal(data, testType)
	if err != nil {
		return nil, err
	}
	return testType, nil
}
func StepOne(event any) (bool, error) {
	return false, nil
}
func StepTwo(event any) (bool, error) {
	return true, nil
}

func TestLoadChain(t *testing.T) {
	LoadChain("test_type", NewTestType, StepOne, StepTwo)
	testEvent := TestType{
		FirstName: "john",
		LastName:  "doe",
	}
	eventData, err := json.Marshal(testEvent)
	assert.Nil(t, err, "shouldn't be an error")
	pChain, ok := GetChain("test_type", eventData)
	assert.Equal(t, ok, true, "chain should exist")
	err = pChain.Process()
	assert.Nil(t, err, "shouldn't be an error")
}
