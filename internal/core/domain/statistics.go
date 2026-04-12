package domain

import (
	"fmt"
	"time"
)

type Statistics struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistics(
	tasksCreated int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}

func (s Statistics) Validate() error {
	if s.TasksCreated < 0 {
		return fmt.Errorf("total tasks cannot be negative")
	}

	if s.TasksCompleted < 0 {
		return fmt.Errorf("completed tasks cannot be negative")
	}

	return nil
}
