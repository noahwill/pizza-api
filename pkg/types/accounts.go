package types

// Account : represents a ZaRnnr account
type Account struct {
	Active      bool     `dynamo:"Active,omitempty" json:"active,omitempty"`
	Address     Address  `dynamo:"Address,omitempty" json:"address,omitempty"`
	CreatedAt   int64    `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Email       string   `dynamo:"Email,omitempty" json:"email,omitempty"`
	FirstName   string   `dynamo:"FirstName,omitempty" json:"firstName,omitempty"`
	LastName    string   `dynamo:"LastName,omitempty" json:"lastName,omitempty"`
	LastUpdated int64    `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Orders      []string `dynamo:"Orders,omitempty" json:"orders,omitempty"`
	Password    string   `dynamo:"Password,omitempty" json:"password,omitempty"` // TODO: this pizza shop dev hasn't done much work with security - pws are only strings
	UUID        string   `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
}

// GetAccountsOutput : ouput for the GetAccountsRoute
type GetAccountsOutput struct {
	Accounts []Account `json:"accounts,omitempty"`
	Error    string    `json:"error,omitempty"`
	Ok       bool      `json:"ok,omitempty"`
}

// GetAccountOutput : output for the GetAccountRoute
type GetAccountOutput struct {
	Account Account `json:"account,omitempty"`
	Error   string  `json:"error,omitempty"`
	Ok      bool    `json:"ok,omitempty"`
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
	Account Account `json:"account,omitempty"`
	Error   string  `json:"error,omitempty"`
	Ok      bool    `json:"ok,omitempty"`
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
	Account Account `json:"account,omitempty"`
	Error   string  `json:"error,omitempty"`
	Ok      bool    `json:"ok,omitempty"`
}

// DeleteAccountOutput : output for the DeleteAccountRoute
type DeleteAccountOutput struct {
	Account Account `json:"account,omitempty"`
	Error   string  `json:"error,omitempty"`
	Ok      bool    `json:"ok,omitempty"`
}
