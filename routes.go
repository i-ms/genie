package genie

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (g *Genie) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	// Recover from panics and log the error
	if g.Debug {
		mux.Use(middleware.Recoverer)
	}
	mux.Use(g.SessionLoad)

	return mux
}
