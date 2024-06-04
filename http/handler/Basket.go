package handler

import (
	"basket/internal/model"
	"basket/internal/repository"
	"encoding/json"
	"github.com/labstack/echo/v4"
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
	userId, ok := c.Get("userId").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User ID missing or invalid in the request")
	}
	basketList, err := h.repo.Get(nil, &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve baskets: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{"basketList": basketList})
}

func (h *BasketHandler) BasketAdd(c echo.Context) error {
	userId, ok := c.Get("userId").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User ID missing or invalid in the request")
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var basket model.Basket
	if err := json.Unmarshal(body, &basket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	basket.UserId = userId // Ensure the basket is linked to the user
	if err := h.repo.Create(&basket); err != nil {
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
	userId, ok := c.Get("userId").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User ID missing or invalid in the request")
	}
	uid := uint(id)

	baskets, err := h.repo.Get(&uid, &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if len(baskets) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"error": "Basket not found"})
	}

	firstBasket := baskets[0]
	if firstBasket.State == model.StateCompleted {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Cannot update a completed basket"})
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var updatedBasket model.Basket
	if err := json.Unmarshal(body, &updatedBasket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedBasket.ID = firstBasket.ID // Maintain the same ID
	if err := h.repo.Update(&updatedBasket); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": updatedBasket})
}

func (h *BasketHandler) GetBasket(c echo.Context) error {

	basketIdParam := c.Param("id")
	basketId, err := strconv.Atoi(basketIdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid basket ID"})
	}
	userId, ok := c.Get("userId").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User ID missing or invalid in the request")
	}
	uid := uint(basketId)
	baskets, err := h.repo.Get(&uid, &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to retrieve basket: "+err.Error())
	}
	if len(baskets) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "No basket found")
	}
	firstBasket := baskets[0]

	return c.JSON(http.StatusOK, map[string]any{"data": firstBasket})
}

func (h *BasketHandler) DeleteBasket(c echo.Context) error {
	basketId := c.Param("id")
	id, err := strconv.Atoi(basketId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	userId, ok := c.Get("userId").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User ID missing or invalid in the request")
	}
	uid := uint(id)

	baskets, err := h.repo.Get(&uid, &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if len(baskets) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"error": "Basket not found"})
	}

	if err := h.repo.Delete(baskets[0]); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"success": "true"})
}
