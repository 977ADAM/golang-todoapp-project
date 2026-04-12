package statisticsservice

import (
	"context"
	"time"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
)

type StatisticsService struct {
	StatisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

func NewStatisticsService(
	statisticsRepository StatisticsRepository,
) *StatisticsService {
	return &StatisticsService{
		StatisticsRepository: statisticsRepository,
	}
}
