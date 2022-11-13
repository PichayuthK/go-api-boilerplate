package service

import (
	"errors"
	"testing"

	repository "github.com/PichayuthK/go-api-boilerplate/business/repository/todo"
	"github.com/stretchr/testify/assert"
)

func Test_GetTodos(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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
	})

	t.Run("fail", func(t *testing.T) {
		// arrange
		todoRepo := repository.NewTodoRepositoryMock()
		todoRepo.On("ListTodo").Return([]repository.Todo{}, errors.New(""))

		todoService := NewTodoService(todoRepo)

		// act
		_, err := todoService.GetTodos()
		assert.Error(t, err)
	})
}

func Test_GetTodoByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		todoId := 1
		todoRepo := repository.NewTodoRepositoryMock()
		todoRepo.On("GetTodo", todoId).Return(&repository.Todo{
			ID:    todoId,
			Title: "test mock",
		}, nil)

		todoService := NewTodoService(todoRepo)

		// act
		todo, err := todoService.GetTodoByID(todoId)

		// assert
		assert.NoError(t, err)
		assert.NotNil(t, todo)
	})

	t.Run("fail", func(t *testing.T) {
		// arrange
		todoId := 1
		todoRepo := repository.NewTodoRepositoryMock()
		todoRepo.On("GetTodo", todoId).Return(&repository.Todo{}, errors.New(""))

		todoService := NewTodoService(todoRepo)

		// act
		todo, err := todoService.GetTodoByID(todoId)

		// assert
		assert.Error(t, err)
		assert.Nil(t, todo)
	})
}
