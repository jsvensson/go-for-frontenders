package api

import "net/http"

type route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

func (h Handler) createRoutes() []route {
	return []route{
		{
			Path:        "/hello",
			Method:      http.MethodGet,
			HandlerFunc: handleHello,
		},
		{
			Path:        "/todo",
			Method:      http.MethodGet,
			HandlerFunc: h.getAllTodos,
		},
		{
			Path:        "/todo",
			Method:      http.MethodPost,
			HandlerFunc: h.createTodo,
		},
		{
			Path:        "/todo/{id}",
			Method:      http.MethodGet,
			HandlerFunc: h.getTodo,
		},
		{
			Path:        "/todo/{id}",
			Method:      http.MethodDelete,
			HandlerFunc: h.deleteTodo,
		},
	}
}
