package userstransporthttp

import (
	"net/http"

	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corehttpresponse "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/response"
	corehttputils "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/utils"
)


type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
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

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)

		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}