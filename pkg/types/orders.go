package types

// Order : represents a ZaRnnr order
type Order struct {
	Active      bool     `dynamo:"Active,omitempty" json:"active,omitempty"`
	AccountID   string   `dynamo:"AccountID,omitempty" json:"accountID,omitempty"`
	Address     Address  `dynamo:"Address,omitempty" json:"address,omitempty"`
	CreatedAt   int64    `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Price       float64  `dynamo:"Cost,omitempty" json:"cost,omitempty"`
	Delivery    bool     `dynamo:"Delivery,omitempty" json:"delivery,omitempty"`
	LastUpdated int64    `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Size        string   `dynamo:"Size,omitempty" json:"size,omitempty"`
	Status      Status   `dynamo:"Status,omitempty" json:"status,omitempty"`
	Toppings    Toppings `dynamo:"Toppings,omitempty" json:"toppings,omitempty"`
	UUID        string   `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
}

// GetOrdersInput : input for the GetOrdersRoute
type GetOrdersInput struct {
	Active *bool `json:"active,omitempty"`
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
	Address  *Address  `json:"address,omitempty"`
	Delivery *bool     `json:"delivery,omitempty"`
	Size     *string   `json:"size,omitempty"`
	Toppings *Toppings `json:"toppings,omitempty"`
}

// CreateOrderOutput : output for the CreateOrderRoute
type CreateOrderOutput struct {
	Order Order  `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}

// UpdateOrderInput : input for the UpdateOrderRoute
type UpdateOrderInput struct {
	Active   *bool     `json:"active,omitempty"`
	Address  *Address  `json:"address,omitempty"`
	Delivery *bool     `json:"delivery,omitempty"`
	Size     *string   `json:"size,omitempty"`
	Status   *string   `json:"status,omitempty"`
	Toppings *Toppings `json:"toppings,omitempty"`
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
