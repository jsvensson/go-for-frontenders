package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsvensson/go-for-frontenders/api"
	"github.com/jsvensson/go-for-frontenders/repository/todo/inmem"
)

func main() {
	router := mux.NewRouter()
	todos := inmem.NewRepository()

	router = api.NewHandler(router, todos)
	_ = http.ListenAndServe(":8080", router)
}
