package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kounosukexxx/taqque/internal/domain/model"
	"github.com/kounosukexxx/taqque/internal/domain/repositories"
)

func (u *TaskUsecase) ListTasks(ctx context.Context) ([]*model.Task, error) {
	tasks, err := u.taskRepository.GetMultiOrderByPriorityDescAndSortKeyAsc(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("taskRepository.GetMultiOrderByPriorityDescAndSortKeyAsc failed: %w", err)
	}
	return tasks, nil
}

func (u *TaskUsecase) PushTask(ctx context.Context, title string, priority float64) ([]*model.Task, error) {
	task, err := model.NewTask(title, priority)
	if err != nil {
		return nil, err
	}

	if err := u.taskRepository.Create(ctx, task); err != nil {
		return nil, fmt.Errorf("taskRepository.CreateTask failed: %w", err)
	}

	tasks, err := u.taskRepository.GetMultiOrderByPriorityDescAndSortKeyAsc(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("taskRepository.GetMultiOrderByPriorityDescAndSortKeyAsc failed: %w", err)
	}
	return tasks, nil
}

func (u *TaskUsecase) PopTask(ctx context.Context, priority float64, baseTime time.Time) ([]*model.Task, error) {
	task, err := u.taskRepository.GetFirstByPriorityOrderBySortKeyAsc(ctx, priority)
	if err != nil {
		if errors.Is(err, repositories.ErrTaskNotFound) {
			return nil, fmt.Errorf("task not found. specified priority: %f", priority)
		}
		return nil, fmt.Errorf("taskRepository.GetFirstByPriorityOrderBySortKeyAsc failed: %w", err)
	}

	task.UpdateAsDeleted(baseTime)
	if err := u.taskRepository.Update(ctx, task); err != nil {
		return nil, fmt.Errorf("taskRepository.CreateTask failed: %w", err)
	}

	tasks, err := u.taskRepository.GetMultiOrderByPriorityDescAndSortKeyAsc(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("taskRepository.GetMultiOrderByPriorityDescAndSortKeyAsc failed: %w", err)
	}
	return tasks, nil
}
