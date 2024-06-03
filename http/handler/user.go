package handler

import (
	"basket/internal/repository"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo repository.User
}

func NewUserHandler(repo repository.User) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Register(c echo.Context) error {
	return nil
}

func (h *UserHandler) Login(c echo.Context) error {
	return nil
}
