package taskspostgresrepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
	coreerrors "github.com/977ADAM/golang-todoapp-project/internal/core/errors"
	corepostgrespool "github.com/977ADAM/golang-todoapp-project/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) GetTask(
	ctx context.Context,
	id int,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	WHERE id=$1;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		id,
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
		if errors.Is(err, corepostgrespool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%d': %w",
				id,
				coreerrors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := taskDomainFromModel(taskModel)

	return taskDomain, nil
}
