package usecase

import "github.com/kounosukexxx/taqque/internal/domain/repositories"

type TaskUsecase struct {
	taskRepository repositories.TaskRepository
}

func NewTaskUsecase(taskRepository repositories.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		taskRepository: taskRepository,
	}
}
