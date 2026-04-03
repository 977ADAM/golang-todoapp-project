package taskspostgresrepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
	coreerrors "github.com/977ADAM/golang-todoapp-project/internal/core/errors"
	corepostgrespool "github.com/977ADAM/golang-todoapp-project/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.tasks (title, description, completed, created_at, completed_at, author_user_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.AuthorUserID,
	)

	var taskModel TaskModel
	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserID,
	)
	if err != nil {
		if errors.Is(err, corepostgrespool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf(
				"%v: user id=%d does not exist: %w",
				err,
				task.AuthorUserID,
				coreerrors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf("scan row: %w", err)
	}

	taskDomain := taskDomainFromModel(taskModel)

	return taskDomain, nil
}
