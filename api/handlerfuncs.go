package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jsvensson/go-for-frontenders/model"
	"github.com/jsvensson/go-for-frontenders/repository/todo"
)

func handleHello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func (h Handler) createTodo(w http.ResponseWriter, r *http.Request) {
	// Limit request body to 100kB
	b := http.MaxBytesReader(w, r.Body, 100_000)
	defer b.Close()

	body, err := io.ReadAll(b)
	if err != nil {
		log.Println("error reading body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var t model.Todo
	if err := json.Unmarshal(body, &t); err != nil {
		log.Println("error unmarshaling JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdTodo, err := h.todos.Create(t.Title, t.Body)
	if err != nil {
		log.Println("error creating todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("created todo with id:%d", createdTodo.ID)

	resp, err := json.Marshal(createdTodo)
	if err != nil {
		log.Println("error marshaling created todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (h Handler) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error parsing ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := h.todos.Get(id)
	if errors.Is(err, todo.ErrNotFound) {
		log.Printf("todo id:%d not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("error getting todo: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(t)
	if err != nil {
		log.Printf("error marshaling todo: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (h Handler) getAllTodos(w http.ResponseWriter, _ *http.Request) {
	t, err := h.todos.GetAll()
	if err != nil {
		log.Printf("error getting todo: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(t)
	if err != nil {
		log.Printf("error marshaling todo: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (h Handler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error parsing ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.todos.Delete(id); err != nil {
		log.Println("error deleting todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("deleted todo with id:%d", id)
}
