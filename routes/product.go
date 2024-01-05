package routes

import (
	"lapakped/handlers"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	e.GET("/products", handlers.FindProducts)
}
