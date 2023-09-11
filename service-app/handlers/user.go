package handlers

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"service-app/core/user"
	"service-app/web"
)

func (h *handler) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.Key).(*web.Values)
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}

	var nu user.NewUser

	err := web.Decode(r, &nu)
	if err != nil {
		h.log.Error(err.Error(), zap.Any("Trace Id", v.TraceId))
		return web.NewRequestError(errors.New(http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
	}

	usr, err := h.us.CreateUser(ctx, nu, v.Now)
	if err != nil {
		return fmt.Errorf("user signup problem: %w", err)
	}
	return web.Respond(ctx, w, usr, http.StatusOK)

}

func (h *handler) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	v, ok := ctx.Value(web.Key).(*web.Values)
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}

	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := web.Decode(r, &login)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	claims, err := h.us.Authenticate(ctx, login.Email, login.Password, v.Now)

	if err != nil {
		h.log.Error(err.Error(), zap.Any("Trace Id", v.TraceId))
		return web.NewRequestError(errors.New("invalid email or password"), http.StatusUnauthorized)
	}

	var tkn struct {
		Token string `json:"token"`
	}

	tkn.Token, err = h.GenerateToken(claims)

	if err != nil {

		return fmt.Errorf("generating token %w", err)
	}

	h.SetCookie(w, tkn.Token)
	return web.Respond(ctx, w, tkn, http.StatusOK)

}

func (h *handler) SetCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
