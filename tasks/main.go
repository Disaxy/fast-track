package main

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrTaskNotFound        = errors.New("task not found")
	ErrTaskIsAlreadyExists = errors.New("task is already exists")
)

type TaskStorage interface {
	GetTask(id uuid.UUID) (*Task, error)
	AddTask(task *Task) error
	RemoveTask(id uuid.UUID) error
}

type Task struct {
	ID        uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (task *Task) SetText(text string) {
	task.Text = text
	task.UpdatedAt = time.Now()
}

func NewTask(text string) *Task {
	task := Task{
		ID:        uuid.New(),
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &task
}

type MapRepo map[uuid.UUID]*Task

func (repo MapRepo) GetTask(id uuid.UUID) (*Task, error) {
	task, ok := repo[id]
	if !ok {
		return nil, ErrTaskNotFound
	}

	return task, nil
}
func (repo MapRepo) AddTask(task *Task) error {
	_, err := repo.GetTask(task.ID)
	if err == nil {
		return ErrTaskIsAlreadyExists
	}

	repo[task.ID] = task

	return nil
}
func (repo MapRepo) RemoveTask(id uuid.UUID) error {
	_, err := repo.GetTask(id)
	if err != nil {
		return err
	}

	delete(repo, id)

	return nil
}

func NewMapRepo() *MapRepo {
	repo := make(MapRepo)
	return &repo
}

type SliceRepo struct {
	Tasks []Task
}

func (repo *SliceRepo) GetTask(id uuid.UUID) (*Task, error) {
	for i := range repo.Tasks {
		if repo.Tasks[i].ID == id {
			return &repo.Tasks[i], nil
		}
	}

	return nil, ErrTaskNotFound
}

func (repo *SliceRepo) AddTask(task *Task) error {
	_, err := repo.GetTask(task.ID)
	if err == nil {
		return ErrTaskIsAlreadyExists
	}

	repo.Tasks = append(repo.Tasks, *task)

	return nil
}

func (repo *SliceRepo) RemoveTask(id uuid.UUID) error {
	for i, task := range repo.Tasks {
		if task.ID == id {
			repo.Tasks = append(repo.Tasks[:i], repo.Tasks[i+1:]...)
			return nil
		}
	}
	return ErrTaskNotFound
}

func NewSliceRepo() *SliceRepo {
	return &SliceRepo{}
}

func main() {}
