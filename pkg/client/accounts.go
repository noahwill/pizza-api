package client

import (
	"encoding/json"
	"pizza-api/pkg/internal"
	"pizza-api/pkg/types"
)

// GetAccounts : gets all accounts that are active
func GetAccounts() (types.GetAccountsOutput, error) {
	out := types.GetAccountsOutput{}

	path := "/api/v1/account"
	result, err := internal.Request("GET", path, nil)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// GetAccount : gets an account for accountID
func GetAccount(accountID string) (types.GetAccountOutput, error) {
	out := types.GetAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("GET", path, nil)
	if err != nil {
		return out, err
	}
	json.Unmarshal(result, &out)
	return out, nil
}

// CreateAccount : creates an account with the given inputs
func CreateAccount(address types.Address, email string, firstName string, lastName string, password string) (types.CreateAccountOutput, error) {
	in := types.CreateAccountInput{
		Address:   &address,
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
		Password:  &password,
	}
	out := types.CreateAccountOutput{}

	path := "/api/v1/account"
	result, err := internal.Request("POST", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccount : updates an account for accountID with the given inputs
func UpdateAccount(accountID string, active bool, address types.Address, email string, firstName string, lastName string, order string, password string) (types.UpdateAccountOutput, error) {
	in := types.UpdateAccountInput{
		Active:    &active,
		Address:   &address,
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
		Order:     &order,
		Password:  &password,
	}
	out := types.UpdateAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountActive : updates the active field for the account of accountID
func UpdateAccountActive(accountID string, active bool) (types.UpdateAccountOutput, error) {
	in := types.UpdateAccountInput{
		Active: &active,
	}
	out := types.UpdateAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountAddress : updates the address for the account of accountID
func UpdateAccountAddress(accountID string, address types.Address) (types.UpdateAccountOutput, error) {
	in := types.UpdateAccountInput{
		Address: &address,
	}
	out := types.UpdateAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountFirstName : updates the first name for the account of accountID
func UpdateAccountFirstName(accountID string, firstName string) (types.UpdateAccountOutput, error) {
	in := types.UpdateAccountInput{
		FirstName: &firstName,
	}
	out := types.UpdateAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountLastName : updates the last name for the account of accountID
func UpdateAccountLastName(accountID string, lastName string) (types.UpdateAccountOutput, error) {
	in := types.UpdateAccountInput{
		LastName: &lastName,
	}
	out := types.UpdateAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrders : appends an order ID on to the orders for the account of accountID
func UpdateAccountOrders(accountID string, order string) (types.UpdateAccountOutput, error) {
	in := types.UpdateAccountInput{
		Order: &order,
	}
	out := types.UpdateAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountPassword : updates the password for the account of accountID
func UpdateAccountPassword(accountID string, password string) (types.UpdateAccountOutput, error) {
	in := types.UpdateAccountInput{
		Password: &password,
	}
	out := types.UpdateAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// DeleteAccount : deletes the  of accountID
func DeleteAccount(accountID string) (types.DeleteAccountOutput, error) {
	out := types.DeleteAccountOutput{}

	path := "/api/v1/account/" + accountID
	result, err := internal.Request("DELETE", path, nil)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}
