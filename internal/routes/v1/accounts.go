package internal

import (
	"fmt"
	"net/http"
	"pizza-api/internal/helpers"
	"pizza-api/pkg/types"
	"pizza-api/utils"

	"github.com/labstack/echo/v4"
)

// GetAccountsRoute : get all active accounts
func GetAccountsRoute(c echo.Context) error {
	out := types.GetAccountsOutput{}
	return c.JSON(http.StatusOK, &out)
}

// GetAccountRoute : get a single account
func GetAccountRoute(c echo.Context) error {
	out := types.GetAccountOutput{}

	return c.JSON(http.StatusOK, &out)
}

// CreateAccountRoute : create an account
func CreateAccountRoute(c echo.Context) error {
	in := types.CreateAccountInput{}
	out := types.CreateAccountOutput{}

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	account, err := helpers.ValidateCreateAccountInput(&in)
	if err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	out.Account = *account

	if err := utils.Config.AccountsTableConn.Put(account).Run(); err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	return c.JSON(http.StatusOK, &out)
}

// UpdateAccountRoute : update an account
func UpdateAccountRoute(c echo.Context) error {
	out := types.UpdateAccountOutput{}
	return c.JSON(http.StatusOK, &out)
}

// DeleteAccountRoute : delete an account
func DeleteAccountRoute(c echo.Context) error {
	out := types.DeleteAccountOutput{}
	return c.JSON(http.StatusOK, &out)
}
