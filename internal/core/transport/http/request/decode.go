package corehttprequest

import (
	"encoding/json"
	"fmt"
	"net/http"

	coreerrors "github.com/977ADAM/golang-todoapp-project/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

// DecodeAndValidateRequest декодирует JSON из тела HTTP-запроса в dest
// и валидирует полученную структуру с помощью тегов validator.
//
// Метод:
//   - читает r.Body и пытается распарсить JSON в dest;
//   - если JSON некорректный, возвращает ошибку с ErrInvalidArgument;
//   - если структура не проходит валидацию, тоже возвращает ErrInvalidArgument.
//
// Ожидается, что dest — это указатель на структуру, например &CreateTaskRequest{}.
//
// Пример использования:
//
//	type CreateTaskRequest struct {
//		Title string `json:"title"`
//		Done  bool   `json:"done"`
//	}
//
//	func (h *Handler) CreateTask(rw http.ResponseWriter, r *http.Request) {
//		var req CreateTaskRequest
//
//		if err := corehttprequest.DecodeAndValidateRequest(r, &req); err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//
//		// req уже заполнен и провалидирован
//		w.WriteHeader(http.StatusCreated)
//	}
func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf(
			"decode json: %v: %w",
			err,
			coreerrors.ErrInvalidArgument,
		)
	}

	var ( err error )

	v, ok := dest.(validatable)
	if ok {
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf(
			"request validation: %v: %w",
			err,
			coreerrors.ErrInvalidArgument,
		)
	}

	return nil
}
