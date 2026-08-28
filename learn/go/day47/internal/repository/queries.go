package repository

const (
	sqlInsertNote = `INSERT INTO notes (title, body) VALUES (?, ?)`
	sqlSelectNote = `SELECT id, title, body FROM notes WHERE id = ?`
	sqlListNotes  = `SELECT id, title, body FROM notes ORDER BY id`
)
