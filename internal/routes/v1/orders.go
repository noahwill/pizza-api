package internal

import (
	"net/http"
	"pizza-api/pkg/types"

	"github.com/labstack/echo/v4"
)

// GetOrdersRoute : get all active Orders
func GetOrdersRoute(c echo.Context) error {
	out := types.GetOrdersOutput{}
	return c.JSON(http.StatusOK, &out)
}

// GetOrderRoute : get a single Order
func GetOrderRoute(c echo.Context) error {
	out := types.GetOrderOutput{}
	return c.JSON(http.StatusOK, &out)
}

// CreateOrderRoute : create an Order
func CreateOrderRoute(c echo.Context) error {
	out := types.CreateOrderOutput{}
	return c.JSON(http.StatusOK, &out)
}

// UpdateOrderRoute : update an Order
func UpdateOrderRoute(c echo.Context) error {
	out := types.UpdateOrderOutput{}
	return c.JSON(http.StatusOK, &out)
}

// DeleteOrderRoute : delete an Order
func DeleteOrderRoute(c echo.Context) error {
	out := types.DeleteOrderOutput{}
	return c.JSON(http.StatusOK, &out)
}
