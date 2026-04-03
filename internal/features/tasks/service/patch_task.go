package tasksservice

import (
	"context"
	"fmt"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
)

func (s *TasksService) PatchTask(
	ctx context.Context,
	id int,
	patch domain.TaskPatch,
) (domain.Task, error) {
	// 1. Достаем из репозитория TASK с этим ID если задача не существует
	// то возвращаем ошибку с пустым юзером.
	task, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from response: %w", err)
	}
	// если получили задачу
	// то применяем патч
	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply task patch: %w", err)
		// если ошибка возрощаем пустую задачу и ошибку
	}

	patchedTask, err := s.tasksRepository.PatchTask(
		ctx,
		id,
		task,
	)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}

	return patchedTask, nil
}
