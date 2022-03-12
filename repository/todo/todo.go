package todo

import (
	"errors"

	"github.com/jsvensson/go-for-frontenders/model"
)

// ErrNotFound is a sentinel error returned when a todo was
// not found in the data store.
var ErrNotFound = errors.New("todo not found")

// A Repository for Todo CRUD functionality.
type Repository interface {
	// Create creates a new todo note with the provided title and body.
	Create(title, body string) (model.Todo, error)

	// Get looks up a todo note by its ID. Returns ErrNotFound
	// if the todo does not exist.
	Get(id int) (model.Todo, error)

	// GetAll returns all todos.
	GetAll() ([]model.Todo, error)

	// Delete deletes a todo with the provided ID.
	Delete(id int) error

	// Exercise: add an Update() for updating an existing entry!
}
