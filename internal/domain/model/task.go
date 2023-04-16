package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// TODO: add validation
type Task struct {
	ID       string
	Title    string
	Priority int

	// defalut value is create time
	SortKey     time.Time
	CreateTime  time.Time
	UpdateTime  time.Time
	DeletedTime *time.Time
}

func NewTask(title string, priority int) (*Task, error) {
	if title == "" {
		return nil, errors.New("title is empty")
	}
	return &Task{
		ID:       uuid.NewString(),
		Title:    title,
		Priority: priority,
	}, nil
}

func (t *Task) UpdateAsDeleted(baseTime time.Time) {
	t.DeletedTime = &baseTime
}
