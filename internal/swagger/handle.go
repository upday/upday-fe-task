package swagger

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"upday.com/upday-task-fe/internal/control"
	"upday.com/upday-task-fe/pkg/model"
)

func handle(w http.ResponseWriter, r *http.Request, methodName string, action func() (interface{}, error)) {
	result, err := action()

	if errors.Is(err, control.ErrNewsNotFound) ||
		errors.Is(err, control.ErrBoardNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if errors.Is(err, model.ErrInvalidAuthor) ||
		errors.Is(err, model.ErrInvalidDescription) ||
		errors.Is(err, model.ErrInvalidImageURL) ||
		errors.Is(err, model.ErrInvalidStateTransition) ||
		errors.Is(err, model.ErrInvalidTitle) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	if err != nil {
		fmt.Println(fmt.Errorf("[ERROR] %s: %w", methodName, err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if result != nil {
		json.NewEncoder(w).Encode(result)
	}
}
