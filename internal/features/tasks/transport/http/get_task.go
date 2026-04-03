package taskstransporthttp

import (
	"net/http"

	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corehttprequest "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/request"
	corehttpresponse "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)

	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, rw)

	taskID, err := corehttprequest.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get taskID path value",
		)

		return
	}

	task, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task",
		)

		return
	}

	response := GetTaskResponse(taskDTOFromDomain(task))

	responseHandler.JSONResponse(response, http.StatusOK)
}
