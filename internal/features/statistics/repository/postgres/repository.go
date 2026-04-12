package statisticspostgresrepository

import corepostgrespool "github.com/977ADAM/golang-todoapp-project/internal/core/repository/postgres/pool"

type StatisticsRepository struct {
	pool corepostgrespool.Pool
}

func NewStatisticsRepository(
	pool corepostgrespool.Pool,
) *StatisticsRepository {
	return &StatisticsRepository{
		pool: pool,
	}
}
