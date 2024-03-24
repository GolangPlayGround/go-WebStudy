package model

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pqHandler struct {
	db *gorm.DB
}

func (s *pqHandler) GetTodos(sessionId string) []*Todo {
	var todos []*Todo
	result := s.db.Where("session_id = ?", sessionId).Find(&todos)
	if result.Error != nil {
		panic(result.Error)
	}
	return todos
}

func (s *pqHandler) AddTodo(name string, sessionId string) *Todo {
	todo := &Todo{
		Name:      name,
		Completed: false,
		CreatedAt: time.Now(),
	}
	result := s.db.Create(todo)
	if result.Error != nil {
		panic(result.Error)
	}
	return todo
}

func (s *pqHandler) RemoveTodo(id int) bool {
	result := s.db.Delete(&Todo{}, id)
	return result.RowsAffected > 0
}
func (s *pqHandler) CompleteTodo(id int, complete bool) bool {
	result := s.db.Model(&Todo{}).Where("id = ?", id).Update("completed", complete)
	return result.RowsAffected > 0
}

func (s *pqHandler) Close() {
	db, err := s.db.DB()
	if err != nil {
		panic(err)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
}

func newPQHandler() DBHandler {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Todo{})
	if err != nil {
		panic(err)
	}

	return &pqHandler{db: db}
}
