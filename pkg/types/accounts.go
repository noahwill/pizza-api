package client

import "pizza-api/utils"

// GetAccountsOutput : ouput for the GetAccountsRoute
type GetAccountsOutput struct {
	Accounts []utils.Account `json:"accounts,omitempty"`
	Error    string          `json:"error,omitempty"`
	Ok       bool            `json:"ok,omitempty"`
}

// GetAccountOutput : output for the GetAccountRoute
type GetAccountOutput struct {
	Account utils.Account `json:"account,omitempty"`
	Error   string        `json:"error,omitempty"`
	Ok      bool          `json:"ok,omitempty"`
}

// CreateAccountInput : input for the CreateAccountRoute
type CreateAccountInput struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Password  *string `json:"password,omitempty"`
}

// CreateAccountOutput : output for the CreateAccountRoute
type CreateAccountOutput struct {
	Account utils.Account `json:"account,omitempty"`
	Error   string        `json:"error,omitempty"`
	Ok      bool          `json:"ok,omitempty"`
}

// UpdateAccountInput : input for the UpdateAccountRoute
type UpdateAccountInput struct {
	Active    *bool   `json:"active,omitempty"`
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Password  *string `json:"password,omitempty"`
}

// UpdateAccountOutput : output for the UpdateAccountRoute
type UpdateAccountOutput struct {
	Account utils.Account `json:"account,omitempty"`
	Error   string        `json:"error,omitempty"`
	Ok      bool          `json:"ok,omitempty"`
}

// DeleteAccountOutput : output for the DeleteAccountRoute
type DeleteAccountOutput struct {
	Account utils.Account `json:"account,omitempty"`
	Error   string        `json:"error,omitempty"`
	Ok      bool          `json:"ok,omitempty"`
}
