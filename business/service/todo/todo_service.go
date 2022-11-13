package service

import (
	"time"

	repository "github.com/PichayuthK/go-api-boilerplate/business/repository/todo"
)

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

	todos, err := s.todoRepo.ListTodo()
	if err != nil {
		return nil, err
	}

	todoResponses := []TodoResponse{}
	for _, todo := range todos {
		todoRes := TodoResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			FetchTiem: time.Now(),
		}
		todoResponses = append(todoResponses, todoRes)
	}

	return todoResponses, nil
}

// GetTodoByID is used to get a single todo by its ID
func (s todoService) GetTodoByID(id int) (*TodoResponse, error) {
	todo, err := s.todoRepo.GetTodo(id)
	if err != nil {
		return nil, err
	}

	todoRes := TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		FetchTiem: time.Now(),
	}

	return &todoRes, nil
}
