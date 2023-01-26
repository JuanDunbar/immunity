package api

import (
	"fmt"
	"net/http"
)

type rulesController struct{}

func (rc *rulesController) GetRules(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("not implemented"))
}

func (rc *rulesController) PostRules(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte(fmt.Sprintf("not implemented")))
}
