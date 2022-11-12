package service

import "time"

// TodoResponse is a return struct from TodoService
// It constructs on top of a Todo returning from repository in order to add or remove some fields
type TodoResponse struct {
	ID        int       `json:"todo_id"`
	Title     string    `json:"title"`
	FetchTiem time.Time `json:"fetch_time"`
}

// TodoService is an interface that shows all functionalities that are offered from TodoService
type TodoService interface {
	GetTodos() ([]TodoResponse, error)
	GetTodoByID(id int) (*TodoResponse, error)
}
