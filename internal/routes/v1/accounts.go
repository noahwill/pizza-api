package internal

import (
	"fmt"
	"net/http"
	v "pizza-api/internal/valid"
	"pizza-api/pkg/types"
	"pizza-api/utils"

	"github.com/apex/log"
	"github.com/labstack/echo/v4"
)

// GetAccountsRoute : get all active accounts
//
// TODO : this will not scale well, need buffer for getting all acounts
//
// TODO : create an input type for this route that allows someone to get All or all Active or all Inactive accounts
func GetAccountsRoute(c echo.Context) error {
	out := types.GetAccountsOutput{}

	// Get active accounts out of the DB
	if err := utils.Config.AccountsTableConn.Scan().Filter("$ = ?", "Active", true).All(&out.Accounts); err != nil {
		out.Error = fmt.Sprintf("Could not get active accounts with error: %s", err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	log.Infof("| Successfully got all %d active accounts from accounts table!", len(out.Accounts))
	return c.JSON(http.StatusOK, &out)
}

// GetAccountRoute : get a single account
func GetAccountRoute(c echo.Context) error {
	accountID := c.Param("uuid")
	out := types.GetAccountOutput{}

	// Get the account out of the DB
	if err := utils.Config.AccountsTableConn.Get("UUID", accountID).One(&out.Account); err != nil {
		out.Error = fmt.Sprintf("Could not get account %s with error: %s", accountID, err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Ok = true
	log.Infof("| Successfully got account %s in the accounts table!", out.Account.UUID)
	return c.JSON(http.StatusOK, &out)
}

// CreateAccountRoute : create an account
func CreateAccountRoute(c echo.Context) error {
	in := types.CreateAccountInput{}
	out := types.CreateAccountOutput{}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	account, err := v.ValidateCreateAccountInput(&in)
	if err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Account = *account

	// Put the new account in the DB
	if err := utils.Config.AccountsTableConn.Put(account).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	log.Infof("| Successfully created a new account %s in the accounts table!", out.Account.UUID)
	return c.JSON(http.StatusOK, &out)
}

// UpdateAccountRoute : update an account
func UpdateAccountRoute(c echo.Context) error {
	accountID := c.Param("uuid")
	in := types.UpdateAccountInput{}
	out := types.UpdateAccountOutput{}

	accountToUpdate := types.Account{}
	if err := utils.Config.AccountsTableConn.Get("UUID", accountID).One(&accountToUpdate); err != nil {
		out.Error = fmt.Sprintf("Could not update account %s with error: %s", accountID, err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not update account %s with error: %s", accountID, err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	updatedAccount, err := v.ValidateUpdateAccountInput(&in, &accountToUpdate)
	if err != nil {
		out.Error = fmt.Sprintf("Could not update account %s with error: %s", accountID, err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Account = *updatedAccount

	// Overwrite the account in the DB with the updated account
	if err := utils.Config.AccountsTableConn.Put(updatedAccount).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not create account %s with error: %s", accountID, err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	log.Infof("| Successfully updated new account %s in the accounts table!", out.Account.UUID)
	return c.JSON(http.StatusOK, &out)
}

// DeleteAccountRoute : delete an account
func DeleteAccountRoute(c echo.Context) error {
	accountID := c.Param("uuid")
	out := types.DeleteAccountOutput{}

	// Check to see if the account exists in the DB before deleting; cache it in the output
	if err := utils.Config.AccountsTableConn.Get("UUID", accountID).One(&out.Account); err != nil {
		out.Error = fmt.Sprintf("Could not delete account %s with error: %s", accountID, err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusBadRequest, &out)
	}

	// Delete the account
	if err := utils.Config.AccountsTableConn.Delete("UUID", out.Account.UUID).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not delete account %s with error: %s", accountID, err.Error())
		out.Ok = false
		log.Errorf("| %s", out.Error)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Account.Active = false

	// Delete the orders associated with the account
	for _, order := range out.Account.Orders {
		if err := utils.Config.OrdersTableConn.Delete("UUID", order).Run(); err != nil {
			out.Error = fmt.Sprintf("Could not delete order %s for deleted account %s with error: %s", order, accountID, err.Error())
			out.Ok = false
			log.Errorf("| %s", out.Error)
			return c.JSON(http.StatusInternalServerError, &out)
		}
	}

	out.Ok = true
	log.Infof("| Successfully deleted account %s from the accounts table along with its %d orders in the orders table!", out.Account.UUID, len(out.Account.Orders))
	return c.JSON(http.StatusOK, &out)
}
