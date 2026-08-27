package store

import "github.com/oz-fatma/agentic-ai-developer/learn/go/day35/internal/models"

type TaskStore interface {
	Create(title string) models.Task
	List() []models.Task
	Get(id int) (models.Task, bool)
	Update(id int, title *string, done *bool) (models.Task, bool)
	Delete(id int) bool
}
