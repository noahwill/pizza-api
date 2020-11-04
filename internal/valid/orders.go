package valid

import (
	"errors"
	"fmt"
	"pizza-api/pkg/types"
	"pizza-api/utils"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

// gets the price for a toppings object
func getToppingsPrice(toppings types.Toppings) float64 {
	var total, toppingsPrice float64

	cheesePrice, _ := types.Cheese[toppings.Cheese]
	saucePrice, _ := types.Sauce[toppings.Sauce]
	for _, topping := range toppings.Toppings {
		toppingsPrice += types.Topping[topping]
	}

	total = cheesePrice + saucePrice + toppingsPrice
	return total
}

// validates a size and returns its price
func validateSize(size string) (float64, error) {
	var price float64

	price, ok := types.Size[size]
	if !ok {
		return price, errors.New("Size is invalid")
	}

	return price, nil
}

// validates a toppings object and returns the sum of its prices
func validateToppings(toppings types.Toppings) (types.Toppings, float64, error) {
	var (
		ok                                                          bool
		cleanToppings                                               types.Toppings
		price, cheesePrice, saucePrice, toppingPrice, toppingsPrice float64
	)

	if toppings.Cheese != "" {
		c := strings.ToLower(strings.TrimSpace(toppings.Cheese))
		if cheesePrice, ok = types.Cheese[c]; !ok {
			return cleanToppings, price, fmt.Errorf("Cheese %s is unavailable", toppings.Cheese)
		}
		cleanToppings.Cheese = c
	}

	if toppings.Sauce != "" {
		s := strings.ToLower(strings.TrimSpace(toppings.Sauce))
		if saucePrice, ok = types.Sauce[s]; !ok {
			return cleanToppings, price, fmt.Errorf("Sauce %s is unavailable", toppings.Sauce)
		}
		cleanToppings.Sauce = s
	}

	for _, topping := range toppings.Toppings {
		t := strings.ToLower(strings.TrimSpace(topping))
		if toppingPrice, ok = types.Topping[t]; !ok {
			return cleanToppings, price, fmt.Errorf("Topping %s is unavailable", topping)
		}
		toppingsPrice += toppingPrice
		cleanToppings.Toppings = append(cleanToppings.Toppings, t)
	}

	price = cheesePrice + saucePrice + toppingsPrice

	return cleanToppings, price, nil
}

// GetOrderByID :
func GetOrderByID(orderID string) (*types.Order, error) {
	var order types.Order
	if err := utils.Config.OrdersTableConn.Get("UUID", orderID).One(&order); err != nil {
		return &order, err
	}
	return &order, nil
}

// ValidateCreateAccountOrderInput : validates CreateAccountOrderInput and constructs an account object to create
func ValidateCreateAccountOrderInput(in *types.CreateAccountOrderInput, account *types.Account) (*types.Order, error) {
	var (
		delivery bool
		address  types.Address
		order    types.Order
	)

	// Validate and set Deliver and Address
	if in.Delivery == nil {
		return &order, errors.New("Specify Delivery")
	} else if in.Delivery != nil {
		if *in.Delivery {
			delivery = true
			if in.Address != nil { // use the given address
				a, err := validateAddress(in.Address)
				if err != nil {
					return &order, err
				}
				address = *a
			} else { // use the account address
				address = account.Address
			}
		} else if !*in.Delivery {
			delivery = false
		}
	}

	// Validate and set Size and size price
	if in.Size == nil {
		return &order, errors.New("Specify Size")
	}
	sizePrice, err := validateSize(strings.ToLower(strings.TrimSpace(*in.Size)))
	if err != nil {
		return &order, err
	}

	// Validate and set Toppings and toppings price
	if in.Toppings == nil {
		return &order, errors.New("Specify Toppings")
	}
	toppings, toppingsPrice, err := validateToppings(*in.Toppings)
	if err != nil {
		return &order, err
	}

	uu, _ := uuid.NewV4()
	order = types.Order{
		Active:      true,
		AccountID:   account.UUID,
		CreatedAt:   time.Now().Unix(),
		Price:       sizePrice + toppingsPrice,
		Delivery:    delivery,
		LastUpdated: time.Now().Unix(),
		Size:        strings.ToLower(strings.TrimSpace(*in.Size)),
		Status:      types.Received,
		Toppings:    toppings,
		UUID:        uu.String(),
	}

	if order.Delivery {
		order.Address = address
	}

	return &order, nil
}

// ValidateUpdateAccountOrderInput : validates UpdateOrderInput and updates the given account object accordingly
func ValidateUpdateAccountOrderInput(in *types.UpdateAccountOrderInput, account *types.Account, order *types.Order) (*types.Order, error) {

	// Validate and update Active
	if in.Active != nil {
		order.Active = *in.Active
	}

	// Validate and update Status
	if in.Status != nil {
		s := strings.TrimSpace(*in.Status)
		status, ok := types.Statuses[s]
		if !ok {
			return order, fmt.Errorf("Status %s is invalid", s)
		}
		if status == types.PickedUp || status == types.Delivered {
			order.Active = false
		}
		order.Status = status
	}

	// Validate and update delivery and Address
	if in.Delivery != nil {
		if order.Status != types.PickedUp && order.Status != types.EnRoute && order.Status != types.Delivered {
			order.Delivery = *in.Delivery
			if *in.Delivery {
				if in.Address != nil { // use the given address
					a, err := validateAddress(in.Address)
					if err != nil {
						return order, err
					}
					order.Address = *a
				} else { // use the account address
					order.Address = account.Address
				}
			} else { // clear the address if delivery is set to false
				order.Address = types.Address{}
			}
		} else if in.Address != nil && in.Delivery == nil {
			return order, errors.New("Address was given, but Delivery was not indicated")
		} else {
			return order, fmt.Errorf("Based on the status %s of the order, Delivery option cannont be updated", order.Status)
		}
	}

	// Validate and update Size
	if in.Size != nil {
		if order.Status == types.Received {
			currentSizePrice, _ := types.Size[order.Size]
			sizePrice, err := validateSize(strings.TrimSpace(*in.Size))
			if err != nil {
				return order, err
			}
			order.Price -= currentSizePrice
			order.Price += sizePrice
		} else {
			return order, fmt.Errorf("Based on the status %s of the order, order Size cannont be updated", order.Status)
		}
	}

	// Validate and update Toppings
	if in.Toppings != nil {
		if order.Status == types.Received {
			currentToppingsPrice := getToppingsPrice(order.Toppings)
			cleanedToppings, toppingsPrice, err := validateToppings(*in.Toppings)
			if err != nil {
				return order, err
			}

			order.Price -= currentToppingsPrice
			order.Price += toppingsPrice

			if cleanedToppings.Cheese != "" {
				order.Toppings.Cheese = cleanedToppings.Cheese
			}
			if cleanedToppings.Sauce != "" {
				order.Toppings.Sauce = cleanedToppings.Sauce
			}
			if len(cleanedToppings.Toppings) != 0 { // Later we're going to want to make this diff the current toppings with the input and allow the user to add to the toppings, not just overwrite them
				order.Toppings.Toppings = cleanedToppings.Toppings
			}
		} else {
			return order, fmt.Errorf("Based on the status %s of the order, order Toppings cannont be updated", order.Status)
		}
	}

	order.LastUpdated = time.Now().Unix()
	return order, nil
}
