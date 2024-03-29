package model

import "time"

type Todo struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type DBHandler interface {
	GetTodos(sessionId string) []*Todo
	AddTodo(sessionId string, name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

func NewDBHandler() DBHandler {
	//handler = newMemoryHandler()
	//return newSqliteHandler(dbConn)
	return newPQHandler()
}
