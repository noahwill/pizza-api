package types

// Order : represents a ZaRnnr order
type Order struct {
	Active      bool     `json:"active,omitempty"`
	AccountID   string   `json:"accountID,omitempty"`
	Address     Address  `json:"address,omitempty"`
	CreatedAt   int64    `json:"createdAt,omitempty"`
	Delivery    bool     `json:"delivery,omitempty"`
	LastUpdated int64    `json:"lastUpdated,omitempty"`
	Pickup      bool     `json:"pickup,omitempty"`
	Size        Size     `json:"size,omitempty"`
	Status      Status   `json:"status,omitempty"`
	Toppings    Toppings `json:"toppings,omitempty"`
	UUID        string   `json:"UUID,omitempty"`
}

// GetOrdersOutput : ouput for the GetOrdersRoute
type GetOrdersOutput struct {
	Orders []Order `json:"orders,omitempty"`
	Error  string  `json:"error,omitempty"`
	Ok     bool    `json:"ok,omitempty"`
}

// GetOrderOutput : output for the GetOrderRoute
type GetOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}

// CreateOrderInput : input for the CreateOrderRoute
type CreateOrderInput struct {
	Active   *bool   `json:"active,omitempty"`
	Cheese   *Cheese `json:"cheese,omitempty"`
	Sauce    *Sauce  `json:"sauce,omitempty"`
	Size     *Size   `json:"size,omitempty"`
	Toppings *string `json:"toppings,omitempty"`
}

// CreateOrderOutput : output for the CreateOrderRoute
type CreateOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}

// UpdateOrderInput : input for the UpdateOrderRoute
type UpdateOrderInput struct {
	Active   *bool   `json:"active,omitempty"`
	Cheese   *Cheese `json:"cheese,omitempty"`
	Sauce    *Sauce  `json:"sauce,omitempty"`
	Size     *Size   `json:"size,omitempty"`
	Toppings *string `json:"topping,omitempty"`
}

// UpdateOrderOutput : output for the UpdateOrderRoute
type UpdateOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}

// DeleteOrderOutput : output for the DeleteOrderRoute
type DeleteOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}
