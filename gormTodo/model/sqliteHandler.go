package model

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteHandler struct {
	db *gorm.DB
}

func (s *sqliteHandler) GetTodos(sessionId string) []*Todo {
	var todos []*Todo
	result := s.db.Where("session_id = ?", sessionId).Find(&todos)
	if result.Error != nil {
		panic(result.Error)
	}
	return todos
}

func (s *sqliteHandler) AddTodo(name string, sessionId string) *Todo {
	todo := Todo{Name: name, SessionID: sessionId, Completed: false, CreatedAt: time.Now()}
	result := s.db.Create(&todo)
	if result.Error != nil {
		panic(result.Error)
	}
	return &todo
}

func (s *sqliteHandler) RemoveTodo(id int) bool {
	result := s.db.Delete(&Todo{}, id)
	return result.RowsAffected > 0
}

func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	result := s.db.Model(&Todo{}).Where("id = ?", id).Update("completed", complete)
	return result.RowsAffected > 0
}

func (s *sqliteHandler) Close() {
	sqlDB, err := s.db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func newSqliteHandler(filepath string) *sqliteHandler {
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		panic(err)
	}

	return &sqliteHandler{db: db}
}
