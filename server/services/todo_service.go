package services

const (
	completed      = true
	isNotCompleted = false
)

func (db *DB) GetAllTodos() (todos []Todo) {
	db.Find(&todos)
	return todos
}

func (db *DB) GetTodoByID(ID string) (todo Todo) {
	db.Where("ID = ?", ID).Find(&todo)
	return todo
}

func (db *DB) CompleteTodo(todo Todo) Todo {
	todo.IsCompleted = true
	db.Save(&todo)
	return todo
}

func (db *DB) GetAllCompletedTodos() (todos []Todo) {
	db.Where("is_completed  = ? ", completed).Find(&todos)
	return todos
}

func (db *DB) GetAllNotCompletedTodos() (todos []Todo) {
	db.Where("is_completed  =  ?", isNotCompleted).Find(&todos)
	return todos
}

func (db *DB) CreateTodo(todo Todo) Todo {
	db.Save(&todo)
	return todo
}

func (db *DB) DeleteTodo(todo Todo) {
	db.Delete(&todo)
}
