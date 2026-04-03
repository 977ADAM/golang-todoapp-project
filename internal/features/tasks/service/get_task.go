package tasksservice

import (
	"context"
	"fmt"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
)

func (s *TasksService) GetTask(
	ctx context.Context,
	id int,
) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from response: %w", err)
	}

	return task, nil
}
