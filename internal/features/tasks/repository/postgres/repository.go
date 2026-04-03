package taskspostgresrepository

import corepostgrespool "github.com/977ADAM/golang-todoapp-project/internal/core/repository/postgres/pool"

type TasksRepository struct {
	pool corepostgrespool.Pool
}

func NewTasksRepository(
	pool corepostgrespool.Pool,
) *TasksRepository {
	return &TasksRepository{
		pool: pool,
	}
}
