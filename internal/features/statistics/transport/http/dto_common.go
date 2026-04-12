package statisticstransporthttp

import (
	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
)

type StatisticsDTOResponse struct {
	TasksCreated               int      `json:"tasks_created"`
	TasksCompleted             int      `json:"tasks_completed"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time"`
}

func StatisticsDTOFromDomain(stats domain.Statistics) StatisticsDTOResponse {
	var avgTime *string
	if stats.TasksAverageCompletionTime != nil {
		duration := stats.TasksAverageCompletionTime.String()
		avgTime = &duration
	}

	return StatisticsDTOResponse{
		TasksCreated:               stats.TasksCreated,
		TasksCompleted:             stats.TasksCompleted,
		TasksCompletedRate:         stats.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}
