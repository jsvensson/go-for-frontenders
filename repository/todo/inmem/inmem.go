package inmem

import (
	"github.com/jsvensson/go-for-frontenders/model"
	"github.com/jsvensson/go-for-frontenders/repository/todo"
)

type Repository struct {
	// a dictionary to store todo entries, with the todo ID as the key.
	storage map[int]model.Todo

	// used as an internal incremental counter to
	// increase the ID as new todos are created.
	counter int
}

// NewRepository initializes and returns a new in-memory repository.
func NewRepository() *Repository {
	return &Repository{
		storage: make(map[int]model.Todo),
	}
}

func (r *Repository) Create(title, body string) (model.Todo, error) {
	// bump the counter on creation of a new todo
	r.counter++
	id := r.counter

	newTodo := model.Todo{
		ID:        id,
		Title:     title,
		Body:      body,
		Completed: false,
	}

	r.storage[id] = newTodo

	return newTodo, nil
}

func (r *Repository) Get(id int) (model.Todo, error) {
	t, ok := r.storage[id]
	if !ok {
		return model.Todo{}, todo.ErrNotFound
	}

	return t, nil
}

func (r *Repository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	for _, t := range r.storage {
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *Repository) Delete(id int) error {
	delete(r.storage, id)
	return nil
}
