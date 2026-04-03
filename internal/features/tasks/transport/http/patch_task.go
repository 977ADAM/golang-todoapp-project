package taskstransporthttp

import (
	"fmt"
	"net/http"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corehttprequest "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/request"
	corehttpresponse "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/response"
	corehttptypes "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       corehttptypes.Nullable[string] `json:"title"`
	Description corehttptypes.Nullable[string] `json:"description"`
	Completed   corehttptypes.Nullable[bool]   `json:"completed"`
}

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("`Title` can't be NULL")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("`Title` must be between 1 and 100 symbols")
		}

	}

	if r.Description.Set {
		if r.Description.Value == nil {
			return fmt.Errorf("`Description` can't be NULL")
		}

		descriptionLen := len([]rune(*r.Description.Value))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf("`Description` must be between 1 and 1000 symbols")
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("`Completed` can't be NULL")
		}
	}

	return nil
}

type PatchTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
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

	var request PatchTaskRequest
	if err := corehttprequest.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed decode and validation HTTP request",
		)

		return
	}

	taskPatch := taskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch task",
		)

		return
	}

	response := PatchTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}
