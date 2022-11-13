package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

const (
	Probe = "/healtz"
)

func NewK8s(r *mux.Router) {
	r.HandleFunc(Probe, healtz().ServeHTTP).Methods("GET")
}

func healtz() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HEALTZ"))
	}
	return http.HandlerFunc(fn)
}
