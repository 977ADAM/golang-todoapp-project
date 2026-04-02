package corehttprequest

import (
	"fmt"
	"net/http"
	"strconv"

	coreerrors "github.com/977ADAM/golang-todoapp-project/internal/core/errors"
)

/*

GET /users/{id}

*/

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf(
			"no key='%s' in path values: %w",
			key,
			coreerrors.ErrInvalidArgument,
		)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"path value='%s' by key='%s' not a valid integer: %v: %w",
			pathValue,
			key,
			err,
			coreerrors.ErrInvalidArgument,
		)
	}

	return val, nil
}
