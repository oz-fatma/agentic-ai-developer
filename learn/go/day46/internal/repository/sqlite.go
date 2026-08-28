package repository

import (
	"database/sql"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day46/internal/models"
)

type SQLiteUserRepo struct {
	db *sql.DB
}

func NewSQLiteUserRepo(db *sql.DB) *SQLiteUserRepo {
	return &SQLiteUserRepo{db: db}
}

func (r *SQLiteUserRepo) Create(name, email string) (models.User, error) {
	res, err := r.db.Exec(`INSERT INTO users (name, email) VALUES (?, ?)`, name, email)
	if err != nil {
		return models.User{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return models.User{}, err
	}
	return models.User{ID: id, Name: name, Email: email}, nil
}

func (r *SQLiteUserRepo) GetByEmail(email string) (models.User, error) {
	var u models.User
	err := r.db.QueryRow(`SELECT id, name, email FROM users WHERE email = ?`, email).
		Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

func (r *SQLiteUserRepo) List() ([]models.User, error) {
	rows, err := r.db.Query(`SELECT id, name, email FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}
