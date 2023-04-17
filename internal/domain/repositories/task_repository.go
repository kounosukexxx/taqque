package repositories

import (
	"context"
	"errors"

	"github.com/kounosukexxx/taqque/internal/domain/model"
)

var ErrTaskNotFound = errors.New("taqque-internal-domain-repositories-task-repository: task not found")

type TaskRepository interface {
	GetMultiOrderByPriorityDescAndSortKeyAsc(ctx context.Context, limit *int) ([]*model.Task, error)
	Create(ctx context.Context, task *model.Task) error
	GetFirstByPriorityOrderBySortKeyAsc(ctx context.Context, priority float64) (*model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	// UpdateMultiAsDeleted(ctx context.Context, id string) error
}
