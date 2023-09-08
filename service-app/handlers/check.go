package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	status := struct {
		Status string
	}{Status: "ok"}

	return json.NewEncoder(w).Encode(status)

}
