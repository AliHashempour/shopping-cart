package main

import (
	"basket/internal/repository"
	"log"

	"basket/http/handler"
	"basket/internal/database"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	basketHandler := handler.NewBasketHandler(repository.NewBasketRepo(db))
	userHandler := handler.NewUserHandler(repository.NewUserRepo(db))

	e := echo.New()
	baskets := e.Group("/basket")
	{
		baskets.GET("/", basketHandler.BasketList)
		baskets.POST("/", basketHandler.BasketAdd)
		baskets.GET("/:id", basketHandler.GetBasket)
		baskets.PATCH("/:id", basketHandler.UpdateBasket)
		baskets.DELETE("/:id", basketHandler.DeleteBasket)
	}

	auth := e.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
