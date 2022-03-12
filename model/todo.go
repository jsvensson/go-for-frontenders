package model

// A Todo should really be finished as soon as possible.
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}
