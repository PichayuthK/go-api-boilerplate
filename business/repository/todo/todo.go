package repository

// Todo struct represent a Todo model in a scope of database
type Todo struct {
	ID    int    `db:"todo_id"`
	Title string `db:"title"`
}

// TodoRepository list all functionalities the are needed to be supported by TodoRepository
type TodoRepository interface {
	ListTodo() ([]Todo, error)
	GetTodo(id int) (*Todo, error)
}
