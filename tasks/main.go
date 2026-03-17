package main

import (
	"errors"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

var (
	ErrTaskNotFound        = errors.New("task not found")
	ErrTaskIsAlreadyExists = errors.New("task is already exists")
)

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

type Repo map[uuid.UUID]*Task

func (repo Repo) GetTask(id uuid.UUID) (*Task, error) {
	task, ok := repo[id]
	if !ok {
		return nil, ErrTaskNotFound
	}

	return task, nil
}
func (repo Repo) AddTask(task *Task) error {
	_, err := repo.GetTask(task.ID)
	if err == nil {
		return ErrTaskIsAlreadyExists
	}

	repo[task.ID] = task

	return nil
}

func (repo Repo) RemoveTask(id uuid.UUID) error {
	_, err := repo.GetTask(id)
	if err != nil {
		return err
	}

	delete(repo, id)

	return nil
}

func NewRepo() *Repo {
	repo := make(Repo)
	return &repo
}

func main() {
	repo := NewRepo()

	task1 := NewTask("task 1")
	task1.SetText("task 1 override")

	task2 := NewTask("task 2")

	_, err := repo.GetTask(task1.ID)
	if err != nil {
		slog.Info(err.Error())
	}

	err = repo.RemoveTask(task1.ID)
	if err != nil {
		slog.Info(err.Error())
	}

	err = repo.AddTask(task1)
	if err != nil {
		slog.Info(err.Error())
	}

	err = repo.AddTask(task2)
	if err != nil {
		slog.Info(err.Error())
	}

	task, err := repo.GetTask(task1.ID)
	if err != nil {
		slog.Info(err.Error())
	}
	if task != nil {
		slog.Info(task.Text)
	}

	err = repo.RemoveTask(task2.ID)
	if err != nil {
		slog.Info(err.Error())
	}
}
