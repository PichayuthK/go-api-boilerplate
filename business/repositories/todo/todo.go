package repository

type Todo struct {
	ID    int    `db:"todo_id"`
	Title string `db:"title"`
}

type TodoRepository interface {
	ListTodo() ([]Todo, error)
	GetTodo(id int) (*Todo, error)
}
