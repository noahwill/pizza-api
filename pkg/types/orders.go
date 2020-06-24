package types

import "pizza-api/utils"

// GetOrdersOutput : ouput for the GetOrdersRoute
type GetOrdersOutput struct {
	Orders []utils.Order `json:"orders,omitempty"`
	Error  string        `json:"error,omitempty"`
	Ok     bool          `json:"ok,omitempty"`
}

// GetOrderOutput : output for the GetOrderRoute
type GetOrderOutput struct {
	Order utils.Order `json:"order,omitempty"`
	Error string      `json:"error,omitempty"`
	Ok    bool        `json:"ok,omitempty"`
}

// CreateOrderInput : input for the CreateOrderRoute
type CreateOrderInput struct {
	Active   *bool         `json:"active,omitempty"`
	Cheese   *utils.Cheese `json:"cheese,omitempty"`
	Sauce    *utils.Sauce  `json:"sauce,omitempty"`
	Size     *utils.Size   `json:"size,omitempty"`
	Toppings *string       `json:"toppings,omitempty"`
}

// CreateOrderOutput : output for the CreateOrderRoute
type CreateOrderOutput struct {
	Order utils.Order `json:"order,omitempty"`
	Error string      `json:"error,omitempty"`
	Ok    bool        `json:"ok,omitempty"`
}

// UpdateOrderInput : input for the UpdateOrderRoute
type UpdateOrderInput struct {
	Active   *bool         `json:"active,omitempty"`
	Cheese   *utils.Cheese `json:"cheese,omitempty"`
	Sauce    *utils.Sauce  `json:"sauce,omitempty"`
	Size     *utils.Size   `json:"size,omitempty"`
	Toppings *string       `json:"topping,omitempty"`
}

// UpdateOrderOutput : output for the UpdateOrderRoute
type UpdateOrderOutput struct {
	Order utils.Order `json:"order,omitempty"`
	Error string      `json:"error,omitempty"`
	Ok    bool        `json:"ok,omitempty"`
}

// DeleteOrderOutput : output for the DeleteOrderRoute
type DeleteOrderOutput struct {
	Order utils.Order `json:"order,omitempty"`
	Error string      `json:"error,omitempty"`
	Ok    bool        `json:"ok,omitempty"`
}
