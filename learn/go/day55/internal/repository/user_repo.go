package repository

import (
	"database/sql"
)

type User struct {
	ID           int64
	Email        string
	PasswordHash string
	Role         string
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(email, hash, role string) (User, error) {
	res, err := r.db.Exec(`INSERT INTO users (email, password_hash, role) VALUES (?, ?, ?)`, email, hash, role)
	if err != nil {
		return User{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}
	return User{ID: id, Email: email, PasswordHash: hash, Role: role}, nil
}

func (r *UserRepo) GetByEmail(email string) (User, error) {
	var u User
	err := r.db.QueryRow(`SELECT id, email, password_hash, role FROM users WHERE email = ?`, email).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role)
	return u, err
}
