package service

import repository "github.com/PichayuthK/go-api-boilerplate/business/repository/todo"

// todoService is an concrete struct fo TodoService interface
// it contians the implementation of all TodoService interface
type todoService struct {
	todoRepo repository.TodoRepository
}

// NewTodoService is a factory function
func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return todoService{todoRepo: todoRepo}
}

// GetTodos is used to get all todos
func (s todoService) GetTodos() ([]TodoResponse, error) {
	panic("unimplemented")
}

// GetTodoByID is used to get a single todo by its ID
func (s todoService) GetTodoByID(id int) (*TodoResponse, error) {
	panic("unimplemented")
}
