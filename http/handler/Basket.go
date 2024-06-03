package handler

import (
	"basket/internal/model"
	"basket/internal/repository"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
)

type BasketHandler struct {
	repo repository.Basket
}

func NewBasketHandler(repo repository.Basket) *BasketHandler {
	return &BasketHandler{repo: repo}
}

func (h *BasketHandler) BasketList(c echo.Context) error {
	basketList, err := h.repo.Get(nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{"basketList": basketList})
}

func (h *BasketHandler) BasketAdd(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var basket model.Basket
	err = json.Unmarshal(body, &basket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = h.repo.Create(&basket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"data": basket})
}

func (h *BasketHandler) UpdateBasket(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid basket ID"})
	}
	uid := uint(id)
	basket, err := h.repo.Get(&uid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Basket not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if basket[0].State == model.StateCompleted {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Cannot update a completed basket"})
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var updatedBasket model.Basket
	err = json.Unmarshal(body, &updatedBasket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedBasket.ID = uint(id)

	err = h.repo.Update(&updatedBasket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": updatedBasket})
}

func (h *BasketHandler) GetBasket(c echo.Context) error {
	basketId := c.Param("id")
	id, err := strconv.Atoi(basketId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	uid := uint(id)
	basket, err := h.repo.Get(&uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	firstBasket := basket[0]

	return c.JSON(http.StatusOK, map[string]any{"data": firstBasket})
}

func (h *BasketHandler) DeleteBasket(c echo.Context) error {
	basketId := c.Param("id")
	id, err := strconv.Atoi(basketId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var deletedBasket model.Basket
	deletedBasket.ID = uint(id)
	err = h.repo.Delete(&deletedBasket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"success": "true"})
}
