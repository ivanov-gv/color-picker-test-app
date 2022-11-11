package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

const (
	healthProbe = "/healtz"
)

func NewK8s(r *mux.Router) {
	// k8s healtz probe
	r.HandleFunc(healthProbe, healtz().ServeHTTP).Methods("GET")
}

func healtz() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HEALTZ"))
	}
	return http.HandlerFunc(fn)
}
