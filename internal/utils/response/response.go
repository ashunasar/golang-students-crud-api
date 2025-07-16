package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}

const (
	StatusOk    = "Ok"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {

	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func GeneralResponse(data any) Response {

	return Response{
		Status: StatusOk,
		Data:   data,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var b strings.Builder

	for i, err := range errs {
		if i > 0 {
			b.WriteString(", ")
		}

		field := strings.ToLower(err.Field()) // Optional: lowercase field name for user-friendliness

		switch err.Tag() {
		case "required":
			b.WriteString(fmt.Sprintf("field '%s' is required", field))
		case "min":
			b.WriteString(fmt.Sprintf("field '%s' must be at least %s characters", field, err.Param()))
		case "max":
			b.WriteString(fmt.Sprintf("field '%s' must be at most %s characters", field, err.Param()))
		case "email":
			b.WriteString(fmt.Sprintf("field '%s' must be a valid email address", field))
		case "gte":
			b.WriteString(fmt.Sprintf("field '%s' must be greater than or equal to %s", field, err.Param()))
		case "lte":
			b.WriteString(fmt.Sprintf("field '%s' must be less than or equal to %s", field, err.Param()))
		default:
			b.WriteString(fmt.Sprintf("field '%s' failed on '%s' validation", field, err.Tag()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  b.String(),
	}
}
