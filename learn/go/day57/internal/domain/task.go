package domain

type Task struct {
	ID    int64
	Title string
}

type TaskRepository interface {
	Save(task Task) (Task, error)
	List() ([]Task, error)
}
