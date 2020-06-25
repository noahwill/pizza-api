package internal

import (
	"fmt"
	"net/http"
	"pizza-api/internal/helpers"
	"pizza-api/pkg/types"
	"pizza-api/utils"

	"github.com/labstack/echo/v4"
)

// GetAccountOrdersRoute : get Orders for an Account
func GetAccountOrdersRoute(c echo.Context) error {
	accountID := c.Param("account")
	in := types.GetOrdersInput{}
	out := types.GetOrdersOutput{}

	_, err := helpers.GetAccountByID(accountID)
	if err != nil {
		out.Error = fmt.Sprintf("Could not get orders for account %s with error: %s", accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not get orders for account %s with error: %s", accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	if in.Active != nil {
		if *in.Active { // Get all active orders
			if err := utils.Config.OrdersTableConn.Scan().Filter("$ = ? AND $ = ?", "Account", accountID, "Active", true).All(&out.Orders); err != nil {
				out.Error = fmt.Sprintf("Could not get active orders for account %s with error: %s", accountID, err.Error())
				out.Ok = false
				return c.JSON(http.StatusInternalServerError, &out)
			}
		} else if !*in.Active { // Get all inactive orders
			if err := utils.Config.OrdersTableConn.Scan().Filter("$ = ? AND $ = ?", "Account", accountID, "Active", false).All(&out.Orders); err != nil {
				out.Error = fmt.Sprintf("Could not get inactive orders for account %s with error: %s", accountID, err.Error())
				out.Ok = false
				return c.JSON(http.StatusInternalServerError, &out)
			}
		}
	} else { // Get all orders
		if err := utils.Config.OrdersTableConn.Scan().Filter("$ = ?", "Account", accountID).All(&out.Orders); err != nil {
			out.Error = fmt.Sprintf("Could not get orders for account %s with error: %s", accountID, err.Error())
			out.Ok = false
			return c.JSON(http.StatusInternalServerError, &out)
		}
	}

	out.Ok = true
	return c.JSON(http.StatusOK, &out)
}

// GetAccountOrderRoute : get a single Order for an Account
func GetAccountOrderRoute(c echo.Context) error {
	accountID := c.Param("account")
	orderID := c.Param("uuid")
	out := types.GetOrderOutput{}

	_, err := helpers.GetAccountByID(accountID)
	if err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %s", accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	// Get the order out of the DB
	if err := utils.Config.OrdersTableConn.Get("UUID", orderID).Filter("$ = ?", "Account", accountID).One(&out.Order); err != nil {
		out.Error = fmt.Sprintf("Could not get order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Ok = true
	return c.JSON(http.StatusOK, &out)
}

// CreateAccountOrderRoute : create an Order for an Account
func CreateAccountOrderRoute(c echo.Context) error {
	accountID := c.Param("account")
	in := types.CreateOrderInput{}
	out := types.CreateOrderOutput{}

	account, err := helpers.GetAccountByID(accountID)
	if err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %s", accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %s", accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	order, err := helpers.ValidateCreateOrderInput(&in, account)
	if err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %s", accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Order = *order

	// Put the new order in the DB
	if err := utils.Config.OrdersTableConn.Put(order).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %s", accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	// Add the order ID to the account object
	account.Orders = append(account.Orders, order.UUID)
	if err := utils.Config.AccountsTableConn.Put(account).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not update account %s orders list with order ID %s with error: %s", accountID, order.UUID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	return c.JSON(http.StatusOK, &out)
}

// UpdateAccountOrderRoute : update an Order for an Account
func UpdateAccountOrderRoute(c echo.Context) error {
	accountID := c.Param("account")
	orderID := c.Param("uuid")
	in := types.UpdateOrderInput{}
	out := types.UpdateOrderOutput{}

	account, err := helpers.GetAccountByID(accountID)
	if err != nil {
		out.Error = fmt.Sprintf("Could not update order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	orderToUpdate := types.Order{}
	if err := utils.Config.OrdersTableConn.Get("UUID", orderID).One(&orderToUpdate); err != nil {
		out.Error = fmt.Sprintf("Could not update order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not update order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	updatedOrder, err := helpers.ValidateUpdateOrderInput(&in, account, &orderToUpdate)
	if err != nil {
		out.Error = fmt.Sprintf("Could not update order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Order = *updatedOrder

	// Overwrite the order in the DB with the updated order
	if err := utils.Config.OrdersTableConn.Put(updatedOrder).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not create order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	return c.JSON(http.StatusOK, &out)
}

// DeleteAccountOrderRoute : delete an Order for an Account
func DeleteAccountOrderRoute(c echo.Context) error {
	accountID := c.Param("account")
	orderID := c.Param("uuid")
	out := types.DeleteOrderOutput{}

	_, err := helpers.GetAccountByID(accountID)
	if err != nil {
		out.Error = fmt.Sprintf("Could not delete order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	// Check to see if the order exists in the DB before deleting; cache it in the output
	if err := utils.Config.OrdersTableConn.Get("UUID", orderID).One(&out.Order); err != nil {
		out.Error = fmt.Sprintf("Could not delete order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	// Delete the order
	if err := utils.Config.OrdersTableConn.Delete("UUID", out.Order.UUID).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not delete order %s for account %s with error: %s", orderID, accountID, err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	return c.JSON(http.StatusOK, &out)
}
