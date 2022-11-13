package repository

import "github.com/stretchr/testify/mock"

// todoRepositoryMock represent a mock struct of testify mock
type todoRepositoryMock struct {
	mock.Mock
}

// NewTodoRepositoryMock is factory function use to setup todo repository mock and it much return a pointer
// it should implement all method of TodoRepository interface
func NewTodoRepositoryMock() *todoRepositoryMock {
	return &todoRepositoryMock{}
}

// ListTodo implement mock version of this function
func (m *todoRepositoryMock) ListTodo() ([]Todo, error) {
	args := m.Called()

	return args.Get(0).([]Todo), args.Error(1)
}

// GetTodo implment mock version of this function
func (m *todoRepositoryMock) GetTodo(id int) (*Todo, error) {
	args := m.Called(id)

	return args.Get(0).(*Todo), args.Error(1)
}
