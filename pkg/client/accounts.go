package client

import (
	"encoding/json"
	"pizza-api/pkg/helpers"
	"pizza-api/pkg/types"
)

// GetAccounts : gets all accounts that are active
func GetAccounts() (types.GetAccountsOutput, error) {
	var out types.GetAccountsOutput

	path := "/api/v1/account"
	result, err := helpers.Request("GET", path, nil)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// GetAccount : gets an account for accountID
func GetAccount(accountID string) (types.GetAccountOutput, error) {
	var out types.GetAccountOutput

	path := "/api/v1/account/" + accountID
	result, err := helpers.Request("GET", path, nil)
	if err != nil {
		return out, err
	}
	json.Unmarshal(result, &out)
	return out, nil
}

// CreateAccount : creates an account with the given inputs
func CreateAccount(address types.Address, email string, firstName string, lastName string, password string) (types.CreateAccountOutput, error) {
	var out types.CreateAccountOutput
	in := types.CreateAccountInput{
		Address:   &address,
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
		Password:  &password,
	}

	path := "/api/v1/account"
	result, err := helpers.Request("POST", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccount : updates an account for accountID with the given inputs
func UpdateAccount(accountID string, active bool, address types.Address, email string, firstName string, lastName string, password string) (types.UpdateAccountOutput, error) {
	var out types.UpdateAccountOutput
	in := types.UpdateAccountInput{
		Active:    &active,
		Address:   &address,
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
		Password:  &password,
	}

	path := "/api/v1/account/" + accountID
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountActive : updates the active field for the account of accountID
func UpdateAccountActive(accountID string, active bool) (types.UpdateAccountOutput, error) {
	var out types.UpdateAccountOutput
	in := types.UpdateAccountInput{
		Active: &active,
	}

	path := "/api/v1/account/" + accountID
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountAddress : updates the address for the account of accountID
func UpdateAccountAddress(accountID string, address types.Address) (types.UpdateAccountOutput, error) {
	var out types.UpdateAccountOutput
	in := types.UpdateAccountInput{
		Address: &address,
	}

	path := "/api/v1/account/" + accountID
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountFirstName : updates the first name for the account of accountID
func UpdateAccountFirstName(accountID string, firstName string) (types.UpdateAccountOutput, error) {
	var out types.UpdateAccountOutput
	in := types.UpdateAccountInput{
		FirstName: &firstName,
	}

	path := "/api/v1/account/" + accountID
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountLastName : updates the last name for the account of accountID
func UpdateAccountLastName(accountID string, lastName string) (types.UpdateAccountOutput, error) {
	var out types.UpdateAccountOutput
	in := types.UpdateAccountInput{
		LastName: &lastName,
	}

	path := "/api/v1/account/" + accountID
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountPassword : updates the password for the account of accountID
func UpdateAccountPassword(accountID string, password string) (types.UpdateAccountOutput, error) {
	var out types.UpdateAccountOutput
	in := types.UpdateAccountInput{
		Password: &password,
	}

	path := "/api/v1/account/" + accountID
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// DeleteAccount : deletes the account of accountID along with any orders made by that account
func DeleteAccount(accountID string) (types.DeleteAccountOutput, error) {
	var out types.DeleteAccountOutput

	path := "/api/v1/account/" + accountID
	result, err := helpers.Request("DELETE", path, nil)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}
