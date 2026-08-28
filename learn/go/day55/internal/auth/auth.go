package auth

import (
	"database/sql"
	"errors"
	"time"

	jwtauth "github.com/oz-fatma/agentic-ai-developer/learn/go/day52/auth"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailTaken         = errors.New("email already registered")
)

type UserStore interface {
	Create(email, hash, role string) (repository.User, error)
	GetByEmail(email string) (repository.User, error)
}

type Service struct {
	users  UserStore
	secret []byte
}

func NewService(users UserStore, secret []byte) *Service {
	return &Service{users: users, secret: secret}
}

func (s *Service) Register(email, password string) (repository.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return repository.User{}, err
	}
	user, err := s.users.Create(email, string(hash), "user")
	if err != nil {
		return repository.User{}, ErrEmailTaken
	}
	return user, nil
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.users.GetByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrInvalidCredentials
		}
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}
	return jwtauth.CreateToken(user.Email, user.Role, s.secret, time.Hour)
}

func (s *Service) ParseToken(token string) (*jwtauth.Claims, error) {
	return jwtauth.ParseToken(token, s.secret)
}
