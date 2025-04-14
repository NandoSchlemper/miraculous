package handlers

import "net/http"

type IAuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct{}

// Register implements IAuthHandler.
func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewAuthHandler() IAuthHandler {
	return &AuthHandler{}
}
