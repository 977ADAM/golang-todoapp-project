package tasksservice

import (
	"context"
	"fmt"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
)

func (s *TasksService) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task domain: %w", err)
	}

	task, err := s.tasksRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("create task: %w", err)
	}

	return task, nil
}
