package service

import (
	"testing"

	repository "github.com/PichayuthK/go-api-boilerplate/business/repository/todo"
	"github.com/stretchr/testify/assert"
)

func Test_GetTodos(t *testing.T) {
	// arrange
	todoRepo := repository.NewTodoRepositoryMock()
	todoRepo.On("ListTodo").Return([]repository.Todo{
		{ID: 1, Title: "Title from mock"},
	}, nil)

	todoService := NewTodoService(todoRepo)

	// act
	todos, err := todoService.GetTodos()

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, todos)
	assert.Equal(t, 1, len(todos))
}
