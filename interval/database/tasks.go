package database

import (
	"restapi-tasks/interval/models"

	"github.com/jmoiron/sqlx"
)

type TaskStore struct {
	db *sqlx.DB
}

func NewTaskStore(db *sqlx.DB) *TaskStore {
	return &TaskStore{db: db}
}

func (s *TaskStore) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	query := `
	SELECT id, title, description, completed, created_at, updated_at 
	FROM tasks 
	order by created_at desc;`
}
