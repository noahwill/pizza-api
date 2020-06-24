package internal

import (
	"fmt"
	"net/http"
	"pizza-api/internal/helpers"
	"pizza-api/pkg/types"
	"pizza-api/utils"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

// GetAccountsRoute : get all active accounts
func GetAccountsRoute(c echo.Context) error {
	out := types.GetAccountsOutput{}
	db := utils.GetAccountsDB()
	db.Query(`
		SELECT * 
		FROM accounts
	`)
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
	db := utils.GetAccountsDB()

	if err := c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}

	if err := helpers.ValidateCreateAccountInput(&in); err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		return c.JSON(http.StatusBadRequest, &out)
	}

	stmt, err := db.Prepare(`
		INSERT INTO accounts(Active, CreatedAt, Email, FirstName, LastName, LastUpdated, Password, UUID) 
		VALUES(?, ?, ?, ?, ?, ?, ?, ? )
	`)
	if err != nil {
		out.Error = fmt.Sprintf("Could not create account with error: %s", err.Error())
		out.Ok = false
		return c.JSON(http.StatusInternalServerError, &out)
	}
	defer stmt.Close()

	uu, _ := uuid.NewV4()
	account := utils.Account{
		Active:      true,
		CreatedAt:   time.Now().Unix(),
		Email:       *in.Email,
		FirstName:   *in.FirstName,
		LastName:    *in.LastName,
		LastUpdated: time.Now().Unix(),
		Password:    *in.Password,
		UUID:        uu.String(),
	}
	out.Account = account

	_, err = stmt.Exec(account.Active, account.CreatedAt, account.Email, account.FirstName, account.LastName, account.LastUpdated, account.Password, account.UUID)
	if err != nil {
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
