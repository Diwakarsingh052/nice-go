package handlers

import (
	"context"
	"errors"
	"net/http"
)

var ErrCheck = errors.New("check trusted error")

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	status := struct {
		Status string
	}{Status: "ok"}
	_ = status
	//return json.NewEncoder(w).Encode(status)
	//return web.Respond(ctx, w, status, http.StatusOK)
	//return errors.New("some kind of error")
	//return web.NewRequestError(ErrCheck, http.StatusBadRequest)
	panic("something bad or worse")
}
