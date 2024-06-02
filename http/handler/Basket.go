package handler

import (
	"basket/internal/model"
	"basket/internal/repository"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type BasketHandler struct {
	repo repository.Basket
}

func NewBasketHandler(repo repository.Basket) *BasketHandler {
	return &BasketHandler{repo: repo}
}

func (h *BasketHandler) BasketList(c echo.Context) error {
	return nil
}

func (h *BasketHandler) BasketAdd(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	basket := new(model.Basket)
	err = json.Unmarshal(body, basket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = h.repo.Create(basket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"data": basket})
}

func (h *BasketHandler) UpdateBasket(c echo.Context) error {
	return nil
}

func (h *BasketHandler) GetBasket(c echo.Context) error {
	return nil
}

func (h *BasketHandler) DeleteBasket(c echo.Context) error {
	return nil
}
