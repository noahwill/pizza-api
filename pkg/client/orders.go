package client

import (
	"encoding/json"
	"pizza-api/pkg/internal"
	"pizza-api/pkg/types"
)

// GetAccountOrders : gets all orders for accountID that are active (can be true, false, or nil)
func GetAccountOrders(accountID string, active bool) (types.GetAccountOrdersOutput, error) {
	in := types.GetAccountOrdersInput{
		Active: &active,
	}
	out := types.GetAccountOrdersOutput{}

	path := "/api/v1/" + accountID + "/order"
	result, err := internal.Request("GET", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// GetAccountOrder : gets an orders for accountID and order ID
func GetAccountOrder(accountID string, orderID string) (types.GetAccountOrderOutput, error) {
	out := types.GetAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order/" + orderID
	result, err := internal.Request("GET", path, nil)
	if err != nil {
		return out, err
	}
	json.Unmarshal(result, &out)
	return out, nil
}

// GetOrderStatus : gets the status of the order for accountID and orderID
func GetOrderStatus(accountID string, orderID string) (types.Status, error) {
	out := types.GetAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order/" + orderID
	result, err := internal.Request("GET", path, nil)
	if err != nil {
		return out.Order.Status, err
	}
	json.Unmarshal(result, &out)
	return out.Order.Status, nil
}

// CreateAccountOrder : creates an order for accountID with the given inputs
func CreateAccountOrder(accountID string, address types.Address, delivery bool, size string, toppings types.Toppings) (types.CreateAccountOrderOutput, error) {
	in := types.CreateAccountOrderInput{
		Address:  &address,
		Delivery: &delivery,
		Size:     &size,
		Toppings: &toppings,
	}
	out := types.CreateAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order"
	result, err := internal.Request("POST", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrder : updates an order for accountID and orderID with the given inputs
func UpdateAccountOrder(accountID string, orderID string, active bool, address types.Address, delivery bool, size string, status types.Status, toppings types.Toppings) (types.UpdateAccountOrderOutput, error) {
	in := types.UpdateAccountOrderInput{
		Active:   &active,
		Address:  &address,
		Delivery: &delivery,
		Size:     &size,
		Toppings: &toppings,
	}
	out := types.UpdateAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order"
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderActive : updates the active field for the order of accountID and orderID
func UpdateAccountOrderActive(accountID string, orderID string, active bool) (types.UpdateAccountOrderOutput, error) {
	in := types.UpdateAccountOrderInput{
		Active: &active,
	}
	out := types.UpdateAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order"
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderDeliveryAddress : updates the delivery and address fields for the order of accountID and orderID
// if address == nil and delivery == true, the order address will be set to the default account address
func UpdateAccountOrderDeliveryAddress(accountID string, orderID string, address types.Address, delivery bool) (types.UpdateAccountOrderOutput, error) {
	in := types.UpdateAccountOrderInput{
		Address:  &address,
		Delivery: &delivery,
	}
	out := types.UpdateAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order"
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderSize : updates the size of the order of accountID and orderID
func UpdateAccountOrderSize(accountID string, orderID string, size string) (types.UpdateAccountOrderOutput, error) {
	in := types.UpdateAccountOrderInput{
		Size: &size,
	}
	out := types.UpdateAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order"
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderStatus : updates the status of the order of accountID and orderID
func UpdateAccountOrderStatus(accountID string, orderID string, status string) (types.UpdateAccountOrderOutput, error) {
	in := types.UpdateAccountOrderInput{
		Status: &status,
	}
	out := types.UpdateAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order"
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderToppings : updates the toppings of the order of accountID and orderID
func UpdateAccountOrderToppings(accountID string, orderID string, toppings types.Toppings) (types.UpdateAccountOrderOutput, error) {
	in := types.UpdateAccountOrderInput{
		Toppings: &toppings,
	}
	out := types.UpdateAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order"
	result, err := internal.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// DeleteAccountOrder : deletes the order of accountID and orderID
func DeleteAccountOrder(accountID string, orderID string) (types.DeleteAccountOrderOutput, error) {
	out := types.DeleteAccountOrderOutput{}

	path := "/api/v1/" + accountID + "/order/" + orderID
	result, err := internal.Request("DELETE", path, nil)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}
