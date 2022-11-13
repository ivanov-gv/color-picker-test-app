// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/ivanov-gv/color-picker-test-app/pkg/config"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	cfg             config.HTTP
	notify          chan error
	shutdownTimeout time.Duration
}

func New(router *mux.Router, cfg config.HTTP) *Server {
	httpServer := &http.Server{
		Handler: router,
		Addr:    cfg.AppPort,
	}

	s := &Server{
		server: httpServer,
		cfg:    cfg,
		notify: make(chan error, 1),
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
