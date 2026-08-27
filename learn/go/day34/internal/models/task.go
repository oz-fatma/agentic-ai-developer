package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTaskInput struct {
	Title string `json:"title"`
}

type UpdateTaskInput struct {
	Title *string `json:"title,omitempty"`
	Done  *bool   `json:"done,omitempty"`
}
