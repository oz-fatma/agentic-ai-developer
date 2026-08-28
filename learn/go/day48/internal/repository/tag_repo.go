package repository

import (
	"database/sql"
)

type Tag struct {
	ID   int64
	Name string
}

type TagRepo struct {
	db *sql.DB
}

func NewTagRepo(db *sql.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (r *TagRepo) Create(name string) (Tag, error) {
	res, err := r.db.Exec(`INSERT INTO tags (name) VALUES (?)`, name)
	if err != nil {
		return Tag{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return Tag{}, err
	}
	return Tag{ID: id, Name: name}, nil
}

func (r *TagRepo) GetByName(name string) (Tag, error) {
	var t Tag
	err := r.db.QueryRow(`SELECT id, name FROM tags WHERE name = ?`, name).Scan(&t.ID, &t.Name)
	return t, err
}

func (r *TagRepo) List() ([]Tag, error) {
	rows, err := r.db.Query(`SELECT id, name FROM tags ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var t Tag
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	return tags, rows.Err()
}
