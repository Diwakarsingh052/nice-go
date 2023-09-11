package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	v, ok := ctx.Value(Key).(*Values) // type assertion
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	v.StatusCode = statusCode
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func RespondError(ctx context.Context, w http.ResponseWriter, err error) error {
	var webErr *Error // first, we need a default value of a custom error type
	//if any layer used it to create an error message or not for external users
	ok := errors.As(err, &webErr)
	if ok {
		appErr := ErrorResponseJson{Error: webErr.err.Error()}
		err := Respond(ctx, w, appErr, webErr.status)
		if err != nil {
			return fmt.Errorf("problem in sending error response %w", err)
		}
		return nil
	}

	//if we can't identify error from a trusted source, we would give a genric msg
	appErr := ErrorResponseJson{Error: http.StatusText(http.StatusInternalServerError)}
	err = Respond(ctx, w, appErr, http.StatusInternalServerError)
	if err != nil {
		return fmt.Errorf("problem in sending error response %w", err)
	}
	return nil

}

func Decode(r *http.Request, val any) error {
	err := json.NewDecoder(r.Body).Decode(&val)
	if err != nil {
		return err
	}
	return nil

}
