package repository

import (
	"database/sql"
)

type Note struct {
	ID    int64
	Title string
	Body  string
}

type NoteRepo struct {
	db *sql.DB
}

func NewNoteRepo(db *sql.DB) *NoteRepo {
	return &NoteRepo{db: db}
}

func (r *NoteRepo) Create(title, body string) (Note, error) {
	res, err := r.db.Exec(sqlInsertNote, title, body)
	if err != nil {
		return Note{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return Note{}, err
	}
	return Note{ID: id, Title: title, Body: body}, nil
}

func (r *NoteRepo) Get(id int64) (Note, error) {
	var n Note
	err := r.db.QueryRow(sqlSelectNote, id).Scan(&n.ID, &n.Title, &n.Body)
	return n, err
}

func (r *NoteRepo) List() ([]Note, error) {
	rows, err := r.db.Query(sqlListNotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Body); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	return notes, rows.Err()
}
