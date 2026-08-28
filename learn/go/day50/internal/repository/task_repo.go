package repository

import (
	"database/sql"
	"errors"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day50/internal/models"
)

var ErrNotFound = errors.New("task not found")

type TaskRepository interface {
	Create(title string, done bool) (models.Task, error)
	Get(id int64) (models.Task, error)
	Update(id int64, title *string, done *bool) (models.Task, error)
	Delete(id int64) error
	List() ([]models.Task, error)
}

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(title string, done bool) (models.Task, error) {
	res, err := r.db.Exec(`INSERT INTO tasks (title, done) VALUES (?, ?)`, title, boolToInt(done))
	if err != nil {
		return models.Task{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{ID: id, Title: title, Done: done}, nil
}

func (r *TaskRepo) Get(id int64) (models.Task, error) {
	var t models.Task
	var done int
	err := r.db.QueryRow(`SELECT id, title, done FROM tasks WHERE id = ?`, id).
		Scan(&t.ID, &t.Title, &done)
	if err == sql.ErrNoRows {
		return models.Task{}, ErrNotFound
	}
	if err != nil {
		return models.Task{}, err
	}
	t.Done = done == 1
	return t, nil
}

func (r *TaskRepo) Update(id int64, title *string, done *bool) (models.Task, error) {
	current, err := r.Get(id)
	if err != nil {
		return models.Task{}, err
	}
	if title != nil {
		current.Title = *title
	}
	if done != nil {
		current.Done = *done
	}
	_, err = r.db.Exec(`UPDATE tasks SET title = ?, done = ? WHERE id = ?`,
		current.Title, boolToInt(current.Done), id)
	if err != nil {
		return models.Task{}, err
	}
	return current, nil
}

func (r *TaskRepo) Delete(id int64) error {
	res, err := r.db.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *TaskRepo) List() ([]models.Task, error) {
	rows, err := r.db.Query(`SELECT id, title, done FROM tasks ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		var done int
		if err := rows.Scan(&t.ID, &t.Title, &done); err != nil {
			return nil, err
		}
		t.Done = done == 1
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
