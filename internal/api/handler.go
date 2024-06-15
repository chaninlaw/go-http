package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chaninlaw/auth/internal/service"
)

type Handler struct {
    userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
    return &Handler{userService: userService}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.userService.FetchAllUsers(context.Background())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Path[len("/user/"):]

    user, err := h.userService.FetchUser(context.Background(), userId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user *service.User

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = h.userService.CreateUser(context.Background(), *user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
