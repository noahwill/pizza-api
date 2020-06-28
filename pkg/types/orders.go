package types

// Order : represents a ZaRnnr order
type Order struct {
	Active      bool     `dynamo:"Active,omitempty"      json:"active,omitempty"`
	AccountID   string   `dynamo:"AccountID,omitempty"   json:"accountID,omitempty"`
	Address     Address  `dynamo:"Address,omitempty"     json:"address,omitempty"`
	CreatedAt   int64    `dynamo:"CreatedAt,omitempty"   json:"createdAt,omitempty"`
	Price       float64  `dynamo:"Cost,omitempty"        json:"cost,omitempty"`
	Delivery    bool     `dynamo:"Delivery,omitempty"    json:"delivery,omitempty"`
	LastUpdated int64    `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Size        string   `dynamo:"Size,omitempty"        json:"size,omitempty"`
	Status      Status   `dynamo:"Status,omitempty"      json:"status,omitempty"`
	Toppings    Toppings `dynamo:"Toppings,omitempty"    json:"toppings,omitempty"`
	UUID        string   `dynamo:"UUID,omitempty"        json:"UUID,omitempty"`
}

// GetAccountOrdersInput : input for the GetOrdersRoute
type GetAccountOrdersInput struct {
	Active *bool `json:"active,omitempty"`
}

// GetAccountOrdersOutput : ouput for the GetOrdersRoute
type GetAccountOrdersOutput struct {
	Orders []Order `json:"orders,omitempty"`
	Error  string  `json:"error,omitempty"`
	Ok     bool    `json:"ok,omitempty"`
}

// GetAccountOrderOutput : output for the GetOrderRoute
type GetAccountOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}

// CreateAccountOrderInput : input for the CreateOrderRoute
type CreateAccountOrderInput struct {
	Address  *Address  `json:"address,omitempty"`
	Delivery *bool     `json:"delivery,omitempty"`
	Size     *string   `json:"size,omitempty"`
	Toppings *Toppings `json:"toppings,omitempty"`
}

// CreateAccountOrderOutput : output for the CreateOrderRoute
type CreateAccountOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}

// UpdateAccountOrderInput : input for the UpdateOrderRoute
type UpdateAccountOrderInput struct {
	Active   *bool     `json:"active,omitempty"`
	Address  *Address  `json:"address,omitempty"`
	Delivery *bool     `json:"delivery,omitempty"`
	Size     *string   `json:"size,omitempty"`
	Status   *string   `json:"status,omitempty"`
	Toppings *Toppings `json:"toppings,omitempty"`
}

// UpdateAccountOrderOutput : output for the UpdateOrderRoute
type UpdateAccountOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}

// DeleteAccountOrderOutput : output for the DeleteOrderRoute
type DeleteAccountOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}
