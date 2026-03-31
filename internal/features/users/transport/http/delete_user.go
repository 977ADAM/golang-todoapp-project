package userstransporthttp

import (
	"net/http"

	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corehttpresponse "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/response"
	corehttputils "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/utils"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)

	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, rw)

	userID, err := corehttputils.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}


	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete user",
		)
		return
	}
	responseHandler.NoContentResponse()
}