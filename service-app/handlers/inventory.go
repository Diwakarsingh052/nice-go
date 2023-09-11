package handlers

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"service-app/auth"
	"service-app/core/inventory"
	"service-app/web"
)

func (h *handler) AddInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.Key).(*web.Values)
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}

	claims, ok := ctx.Value(auth.Key).(auth.Claims)
	if !ok {
		return web.NewRequestError(errors.New("not authenticated"), http.StatusUnauthorized)
	}
	var newInv inventory.NewShirtInventory
	err := web.Decode(r, &newInv)
	if err != nil {
		return err
	}

	inv, err := h.inv.CreateInventory(ctx, newInv, claims.Subject, v.Now)

	if err != nil {
		h.log.Error(err.Error(), zap.Any("Trace Id", v.TraceId))
		return web.NewRequestError(errors.New("problem in creating inventory"), http.StatusBadRequest)
	}

	return web.Respond(ctx, w, inv, http.StatusOK)

}
