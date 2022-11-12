package repository

import "github.com/jmoiron/sqlx"

type todoRepositoryDB struct {
	db *sqlx.DB
}

func NewTodoRepositoryDB(db *sqlx.DB) TodoRepository {
	return todoRepositoryDB{db: db}
}

func (r todoRepositoryDB) ListTodo() ([]Todo, error) {
	panic("unimplemented")
}

func (r todoRepositoryDB) GetTodo(id int) (*Todo, error) {
	panic("unimplemented")
}
