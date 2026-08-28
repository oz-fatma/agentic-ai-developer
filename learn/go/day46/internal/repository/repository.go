package repository

import "github.com/oz-fatma/agentic-ai-developer/learn/go/day46/internal/models"

type UserRepository interface {
	Create(name, email string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	List() ([]models.User, error)
}
