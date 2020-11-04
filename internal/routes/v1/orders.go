package internal

import (
	"fmt"
	"net/http"
	v "pizza-api/internal/valid"
	"pizza-api/pkg/types"
	"pizza-api/utils"
	"sort"
	"time"

	"github.com/apex/log"
	"github.com/labstack/echo/v4"
)

// GetAccountOrdersRoute : get Orders for an Account
func GetAccountOrdersRoute(c echo.Context) error {
	var (
		logText string
		orders  []types.Order
		in      types.GetAccountOrdersInput
		out     types.GetAccountOrdersOutput
	)
	accountID := c.Param("account")

	if _, err := v.GetAccountByID(accountID); err != nil {
		out.Error = fmt.Sprintf("Could not get orders for account %s with error: %v", accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not get orders for account %s with error: %v", accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	if in.Active != nil {
		if *in.Active { // Get all active orders
			logText = " active"
			if err := utils.Config.OrdersTableConn.Scan().Filter("$ = ? AND $ = ?", "AccountID", accountID, "Active", true).All(&orders); err != nil {
				out.Error = fmt.Sprintf("Could not get active orders for account %s with error: %v", accountID, err)
				log.Errorf("| %s", out.Error)
				return c.JSON(http.StatusInternalServerError, &out)
			}
		} else if !*in.Active { // Get all inactive orders
			logText = " inactive"
			if err := utils.Config.OrdersTableConn.Scan().Filter("$ = ? AND $ = ?", "AccountID", accountID, "Active", false).All(&orders); err != nil {
				out.Error = fmt.Sprintf("Could not get inactive orders for account %s with error: %v", accountID, err)
				log.Errorf("| %s", out.Error)
				return c.JSON(http.StatusInternalServerError, &out)
			}
		}
	} else { // Get all orders
		if err := utils.Config.OrdersTableConn.Scan().Filter("$ = ?", "AccountID", accountID).All(&orders); err != nil {
			out.Error = fmt.Sprintf("Could not get orders for account %s with error: %v", accountID, err)
			log.Errorf("| %s", out.Error)
			return c.JSON(http.StatusInternalServerError, &out)
		}
	}

	out.Orders = orders
	out.Ok = true
	log.Infof("| Successfully got all %s%s orders for account %s from the orders table!", len(out.Orders), logText, accountID)
	return c.JSON(http.StatusOK, &out)

}

// GetAccountOrderRoute : get a single Order for an Account
func GetAccountOrderRoute(c echo.Context) error {
	var out types.GetAccountOrderOutput
	accountID := c.Param("account")
	orderID := c.Param("uuid")

	if _, err := v.GetAccountByID(accountID); err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %v", accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	// Get the order out of the DB
	if err := utils.Config.OrdersTableConn.Get("UUID", orderID).Filter("$ = ?", "AccountID", accountID).One(&out.Order); err != nil {
		out.Error = fmt.Sprintf("Could not get order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Ok = true
	log.Infof("| Successfully got order %s for account %s from the orders table!", out.Order.UUID, out.Order.AccountID)
	return c.JSON(http.StatusOK, &out)
}

// CreateAccountOrderRoute : create an Order for an Account
func CreateAccountOrderRoute(c echo.Context) error {
	var (
		account *types.Account
		order   *types.Order
		in      types.CreateAccountOrderInput
		out     types.CreateAccountOrderOutput
	)

	accountID := c.Param("account")

	account, err := v.GetAccountByID(accountID)
	if err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %v", accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %v", accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	order, err = v.ValidateCreateAccountOrderInput(&in, account)
	if err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %v", accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Order = *order

	// Put the new order in the DB
	if err := utils.Config.OrdersTableConn.Put(order).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not create order for account %s with error: %v", accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	// Add the order ID to the account object
	account.Orders = append(account.Orders, order.UUID)
	if err := utils.Config.AccountsTableConn.Put(account).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not update account %s orders list with order ID %s with error: %v", accountID, order.UUID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	log.Infof("| Successfully created order %s for account %s in the orders table!", out.Order.UUID, out.Order.AccountID)
	return c.JSON(http.StatusOK, &out)
}

// UpdateAccountOrderRoute : update an Order for an Account
func UpdateAccountOrderRoute(c echo.Context) error {
	var (
		orderToUpdate types.Order
		updatedOrder  *types.Order
		in            types.UpdateAccountOrderInput
		out           types.UpdateAccountOrderOutput
	)

	accountID := c.Param("account")
	orderID := c.Param("uuid")

	account, err := v.GetAccountByID(accountID)
	if err != nil {
		out.Error = fmt.Sprintf("Could not update order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	if err := utils.Config.OrdersTableConn.Get("UUID", orderID).One(&orderToUpdate); err != nil {
		out.Error = fmt.Sprintf("Could not update order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not update order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	updatedOrder, err = v.ValidateUpdateAccountOrderInput(&in, account, &orderToUpdate)
	if err != nil {
		out.Error = fmt.Sprintf("Could not update order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Order = *updatedOrder

	// Overwrite the order in the DB with the updated order
	if err := utils.Config.OrdersTableConn.Put(updatedOrder).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not create order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	log.Infof("| Successfully updated order %s for account %s in the orders table!", out.Order.UUID, out.Order.AccountID)
	return c.JSON(http.StatusOK, &out)
}

// DeleteAccountOrderRoute : delete an Order for an Account
func DeleteAccountOrderRoute(c echo.Context) error {
	var out types.DeleteAccountOrderOutput

	accountID := c.Param("account")
	orderID := c.Param("uuid")

	account, err := v.GetAccountByID(accountID)
	if err != nil {
		out.Error = fmt.Sprintf("Could not delete order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	// Check to see if the order exists in the DB before deleting; cache it in the output
	if err := utils.Config.OrdersTableConn.Get("UUID", orderID).One(&out.Order); err != nil {
		out.Error = fmt.Sprintf("Could not delete order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	// Delete the order
	if err := utils.Config.OrdersTableConn.Delete("UUID", out.Order.UUID).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not delete order %s for account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	// Remove the order from the account orders list and update the account in the DB
	var a, b []string
	accountOrders := account.Orders
	sort.Strings(accountOrders)
	removeIdx := sort.SearchStrings(accountOrders, out.Order.UUID)
	a = accountOrders[:removeIdx-1]
	b = accountOrders[removeIdx:]
	account.Orders = append(a, b...)
	account.LastUpdated = time.Now().Unix()

	if err := utils.Config.AccountsTableConn.Put(&account).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not remove order %s from list of orders on account %s with error: %v", orderID, accountID, err)
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Order.Active = false
	out.Ok = true
	log.Infof("| Successfully deleted order %s for account %s from the orders table and removed the order ID from the account object in the accounts table!", out.Order.UUID, out.Order.AccountID)
	return c.JSON(http.StatusOK, &out)
}
