package taskstransporthttp

import (
	"net/http"

	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corehttprequest "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/request"
	corehttpresponse "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/response"
)

func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
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

	if err := h.tasksService.DeleteTask(ctx, taskID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete task",
		)
		return
	}
	responseHandler.NoContentResponse()
}
