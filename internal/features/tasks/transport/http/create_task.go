package taskstransporthttp

import (
	"net/http"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corehttprequest "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/request"
	corehttpresponse "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorUserId int     `json:"author_user_id" validate:"required"`
}

type CreateTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, rw)

	var request CreateTaskRequest
	if err := corehttprequest.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	taskDomain := domainFromDTO(request)

	taskDomain, err := h.tasksService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")

		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(
		response,
		http.StatusCreated,
	)
}

func domainFromDTO(dto CreateTaskRequest) domain.Task {
	return domain.NewTaskUninitialized(
		dto.Title,
		dto.Description,
		dto.AuthorUserId,
	)
}
