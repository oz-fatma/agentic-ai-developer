package memory

import (
	"sync"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day57/internal/domain"
)

type TaskRepo struct {
	mu     sync.Mutex
	nextID int64
	items  []domain.Task
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (r *TaskRepo) Save(task domain.Task) (domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nextID++
	task.ID = r.nextID
	r.items = append(r.items, task)
	return task, nil
}

func (r *TaskRepo) List() ([]domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]domain.Task, len(r.items))
	copy(out, r.items)
	return out, nil
}
