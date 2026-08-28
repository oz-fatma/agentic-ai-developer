package usecase

import "github.com/oz-fatma/agentic-ai-developer/learn/go/day57/internal/domain"

type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Add(title string) (domain.Task, error) {
	return s.repo.Save(domain.Task{Title: title})
}

func (s *TaskService) List() ([]domain.Task, error) {
	return s.repo.List()
}
