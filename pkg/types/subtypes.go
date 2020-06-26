package types

// Address : represents a ZaRnnr order address
type Address struct {
	ExtendedAddress string `dynamo:"ExtendedAddress,omitempty" json:"extendedAddress,omitempty"`
	Locality        string `dynamo:"Locality,omitempty" json:"locality,omitempty"`
	PostalCode      string `dynamo:"PostalCode,omitempty" json:"postalCode,omitempty"`
	Region          string `dynamo:"Region,omitempty" json:"region,omitempty"`
	StreetAddress   string `dynamo:"StreetAddress,omitempty" json:"streetAddress,omitempty"`
}

var (
	// Cheese name to price
	Cheese = map[string]float64{
		"cheddar":    3.00,
		"mozzarella": 3.00,
		"parmesan":   4.00,
	}

	// Sauce name to price
	Sauce = map[string]float64{
		"tomato":   2.00,
		"white":    2.00,
		"barbeque": 3.00,
	}

	// Size name to price
	Size = map[string]float64{
		"small":  8.00,
		"medium": 16.00,
		"large":  32.00,
		"party":  48.00,
	}

	// Statuses available
	Statuses = map[string]Status{
		"Received":  Received,
		"Assebly":   Assembly,
		"Baking":    Baking,
		"Ready":     Ready,
		"PickedUp":  PickedUp,
		"EnRoute":   EnRoute,
		"Delivered": Delivered,
	}

	// Topping name to price
	Topping = map[string]float64{
		"anchovies": 0.75,
		"artichoke": 1.25,
		"basil":     0.75,
		"chicken":   1.75,
		"ham":       1.25,
		"kale":      0.75,
		"olives":    0.50,
		"onion":     0.50,
		"pepperoni": 1.25,
		"pineapple": 100.00,
		"tomato":    0.75,
	}
)

// Toppings : represents a ZaRnnr order toppings
type Toppings struct {
	Cheese   string   `dynamo:"Cheese,omitempty" json:"cheese,omitempty"`
	Sauce    string   `dynamo:"Sauce,omitempty" json:"sauce,omitempty"`
	Toppings []string `dynamo:"Toppings,omitempty" json:"toppings,omitempty"`
}

// Status : represents the status of an order
type Status string

const (
	// Received status
	Received Status = "Received"
	// Assembly status
	Assembly Status = "Assembly"
	// Baking status
	Baking Status = "Baking"
	// Ready status
	Ready Status = "Ready"
	// PickedUp status
	PickedUp Status = "PickedUp"
	// EnRoute status
	EnRoute Status = "EnRoute"
	// Delivered status
	Delivered Status = "Delivered"
)
