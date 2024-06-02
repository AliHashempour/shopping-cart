package handler

import (
	"basket/internal/database"
	"basket/internal/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func BasketList(c echo.Context) error {
	return nil
}

func BasketAdd(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var basket model.Basket
	err = json.Unmarshal(body, &basket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	result := database.DB.Create(&basket)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"data": basket})
}

func UpdateBasket(c echo.Context) error {
	return nil
}

func GetBasket(c echo.Context) error {
	return nil
}

func DeleteBasket(c echo.Context) error {
	return nil
}
