package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// struct
type Products struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Category string `json:"category"`
	Rating string `json:"rating"`
}

// dummies
var products = []Products{
	{
		Id: "1",
		Name: "Corkcicle",
		Price: "750000",
		Category: "Utilities",
		Rating: "9",
	},
	{
		Id: "1",
		Name: "Hydroflask",
		Price: "550000",
		Category: "Utilities",
		Rating: "7",
	},
}

// get all products
func FindProducts(c echo.Context) error{
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(products)
}