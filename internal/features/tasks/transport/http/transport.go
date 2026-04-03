package taskstransporthttp

import (
	"context"
	"net/http"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
	corehttpserver "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	tasksService TasksService
}

type TasksService interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
	GetTasks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Task, error)
	GetTask(
		ctx context.Context,
		id int,
	) (domain.Task, error)
	DeleteTask(
		ctx context.Context,
		id int,
	) error
	PatchTask(
		ctx context.Context,
		id int,
		task domain.TaskPatch,
	) (domain.Task, error)
}

func NewTasksHTTPHandler(tasksService TasksService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHTTPHandler) Routes() []corehttpserver.Route {
	return []corehttpserver.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: h.DeleteTask,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/tasks/{id}",
			Handler: h.PatchTask,
		},
	}
}
