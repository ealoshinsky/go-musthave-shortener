package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type app struct {
	mux *chi.Mux
	s   *http.Server
}

func NewRouter() *app {
	return &app{mux: chi.NewRouter(), s: &http.Server{}}
}

func (a *app) Run(addr string) *app {
	a.s.Addr = addr
	if err := http.ListenAndServe(addr, a.mux); err != nil {
		log.Println("Server listen:", err)
	}
	return a
}

func (a *app) Stop(ctx context.Context) {
	a.s.Shutdown(ctx)
}
