package main

import (
	"fmt"
	"lapakped/routes"

	"github.com/labstack/echo/v4"
)



func main() {
	e:= echo.New()

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("server running on localhost: 3342")
	e.Logger.Fatal(e.Start("localhost:3342"))
}

