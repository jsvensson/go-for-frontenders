package api

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/jsvensson/go-for-frontenders/repository/todo"
)

// A Handler contains the API for interacting with todos.
type Handler struct {
	todos todo.Repository
}

// NewHandler creates a new API handler.
func NewHandler(router *mux.Router, todos todo.Repository) *mux.Router {
	h := Handler{
		todos: todos,
	}

	for _, r := range h.createRoutes() {
		log.Printf("registering API handler %s %s", r.Method, r.Path)
		router.Path(r.Path).
			Methods(r.Method).
			HandlerFunc(r.HandlerFunc)
	}

	return router
}
