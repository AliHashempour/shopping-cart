package main

import (
	"fmt"
	"log"

	"basket/api/handler"
	"basket/internal/database"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize the database
	_, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Println("Server started...")

	// Create a new Echo instance
	e := echo.New()
	baskets := e.Group("/basket")
	{
		baskets.GET("/", handler.BasketList)
		baskets.POST("/", handler.BasketAdd)
		baskets.GET("/:id", handler.GetBasket)
		baskets.PATCH("/:id", handler.UpdateBasket)
		baskets.DELETE("/:id", handler.DeleteBasket)
	}

	// Start the server on port 8080
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
