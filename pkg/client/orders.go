package client

import (
	"encoding/json"
	"pizza-api/pkg/helpers"
	"pizza-api/pkg/types"
)

// GetAccountOrders : gets all orders for accountID that are active (can be true, false, or nil)
func GetAccountOrders(accountID string, active bool) (types.GetAccountOrdersOutput, error) {
	var out types.GetAccountOrdersOutput
	in := types.GetAccountOrdersInput{
		Active: &active,
	}

	path := "/api/v1/" + accountID + "/order"
	result, err := helpers.Request("GET", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// GetAccountOrder : gets an orders for accountID and order ID
func GetAccountOrder(accountID string, orderID string) (types.GetAccountOrderOutput, error) {
	var out types.GetAccountOrderOutput

	path := "/api/v1/" + accountID + "/order/" + orderID
	result, err := helpers.Request("GET", path, nil)
	if err != nil {
		return out, err
	}
	json.Unmarshal(result, &out)
	return out, nil
}

// GetOrderStatus : gets the status of the order for accountID and orderID
func GetOrderStatus(accountID string, orderID string) (types.Status, error) {
	var out types.GetAccountOrderOutput

	path := "/api/v1/" + accountID + "/order/" + orderID
	result, err := helpers.Request("GET", path, nil)
	if err != nil {
		return out.Order.Status, err
	}
	json.Unmarshal(result, &out)
	return out.Order.Status, nil
}

// CreateAccountOrder : creates an order for accountID with the given inputs
func CreateAccountOrder(accountID string, address types.Address, delivery bool, size string, toppings types.Toppings) (types.CreateAccountOrderOutput, error) {
	var out types.CreateAccountOrderOutput
	in := types.CreateAccountOrderInput{
		Address:  &address,
		Delivery: &delivery,
		Size:     &size,
		Toppings: &toppings,
	}

	path := "/api/v1/" + accountID + "/order"
	result, err := helpers.Request("POST", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrder : updates an order for accountID and orderID with the given inputs
func UpdateAccountOrder(accountID string, orderID string, active bool, address types.Address, delivery bool, size string, status types.Status, toppings types.Toppings) (types.UpdateAccountOrderOutput, error) {
	var out types.UpdateAccountOrderOutput
	in := types.UpdateAccountOrderInput{
		Active:   &active,
		Address:  &address,
		Delivery: &delivery,
		Size:     &size,
		Toppings: &toppings,
	}

	path := "/api/v1/" + accountID + "/order"
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderActive : updates the active field for the order of accountID and orderID
func UpdateAccountOrderActive(accountID string, orderID string, active bool) (types.UpdateAccountOrderOutput, error) {
	var out types.UpdateAccountOrderOutput
	in := types.UpdateAccountOrderInput{
		Active: &active,
	}

	path := "/api/v1/" + accountID + "/order"
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderDeliveryAddress : updates the delivery and address fields for the order of accountID and orderID
// if address == nil and delivery == true, the order address will be set to the default account address
func UpdateAccountOrderDeliveryAddress(accountID string, orderID string, address types.Address, delivery bool) (types.UpdateAccountOrderOutput, error) {
	var out types.UpdateAccountOrderOutput
	in := types.UpdateAccountOrderInput{
		Address:  &address,
		Delivery: &delivery,
	}

	path := "/api/v1/" + accountID + "/order"
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderSize : updates the size of the order of accountID and orderID
func UpdateAccountOrderSize(accountID string, orderID string, size string) (types.UpdateAccountOrderOutput, error) {
	var out types.UpdateAccountOrderOutput
	in := types.UpdateAccountOrderInput{
		Size: &size,
	}

	path := "/api/v1/" + accountID + "/order"
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderStatus : updates the status of the order of accountID and orderID
func UpdateAccountOrderStatus(accountID string, orderID string, status string) (types.UpdateAccountOrderOutput, error) {
	var out types.UpdateAccountOrderOutput
	in := types.UpdateAccountOrderInput{
		Status: &status,
	}

	path := "/api/v1/" + accountID + "/order"
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// UpdateAccountOrderToppings : updates the toppings of the order of accountID and orderID
func UpdateAccountOrderToppings(accountID string, orderID string, toppings types.Toppings) (types.UpdateAccountOrderOutput, error) {
	var out types.UpdateAccountOrderOutput
	in := types.UpdateAccountOrderInput{
		Toppings: &toppings,
	}

	path := "/api/v1/" + accountID + "/order"
	result, err := helpers.Request("PUT", path, in)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}

// DeleteAccountOrder : deletes the order of accountID and orderID
func DeleteAccountOrder(accountID string, orderID string) (types.DeleteAccountOrderOutput, error) {
	var out types.DeleteAccountOrderOutput

	path := "/api/v1/" + accountID + "/order/" + orderID
	result, err := helpers.Request("DELETE", path, nil)
	if err != nil {
		return out, err
	}

	json.Unmarshal(result, &out)
	return out, nil
}
