package handler

import (
	jwtUtil "basket/http/jwt"
	"basket/internal/model"
	"basket/internal/repository"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

type UserHandler struct {
	repo repository.User
}

func NewUserHandler(repo repository.User) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Register(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to read request body"})
	}

	var user model.User
	if err := json.Unmarshal(body, &user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user data format"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	if err := h.repo.Create(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c echo.Context) error {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(c.Request().Body).Decode(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user, err := h.repo.GetBy("username = ?", credentials.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	token, err := jwtUtil.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}
