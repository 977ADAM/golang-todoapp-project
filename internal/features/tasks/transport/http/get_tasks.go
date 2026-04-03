package taskstransporthttp

import (
	"fmt"
	"net/http"

	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corehttprequest "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/request"
	corehttpresponse "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, rw)

	userID, limit, offset, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'user_id'/'limit'/'offset' query param",
		)

		return
	}

	taskDomains, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed get tasks",
		)
	}

	response := GetTasksResponse(tasksDTOFromDomains(taskDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
		userIDQueryParamKey = "user_id"
	)

	userID, err := corehttprequest.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	limit, err := corehttprequest.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := corehttprequest.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return userID, limit, offset, nil
}
