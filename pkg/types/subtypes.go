package types

// Address : represents a ZaRnnr order address
type Address struct {
	CreatedAt       int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Country         string `dynamo:"Country,omitempty" json:"country,omitempty"`
	ExtendedAddress string `dynamo:"EdtendedAddress,omitempty" json:"extendedAddress,omitempty"`
	LastUpdated     int64  `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Locality        string `dynamo:"Locality,omitempty" json:"locality,omitempty"`
	PostalCode      string `dynamo:"PostalCode,omitempty" json:"postalCode,omitempty"`
	Region          string `dynamo:"Region,omitempty" json:"region,omitempty"`
	StreetAddress   string `dynamo:"StreetAddress,omitempty" json:"startTime,omitempty"`
	UUID            string `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
}

// Cheese : represents an available pizza cheese
type Cheese string

const (
	chCheddar    Cheese = "cheddar"
	chMozzarella Cheese = "mozzarella"
	chParmesan   Cheese = "parmesan"
)

// Toppings : represents a ZaRnnr order toppings
type Toppings struct {
	Cheese   Cheese    `dynamo:"Cheese,omitempty" json:"cheese,omitempty"`
	Sauce    Sauce     `dynamo:"Sauce,omitempty" json:"sauce,omitempty"`
	Toppings []Topping `dynamo:"Toppings,omitempty" json:"toppings,omitempty"`
}

// Sauce : represents an available pizza sauce
type Sauce string

const (
	scTomato   Sauce = "tomato"
	scWhite    Sauce = "white"
	scBarbeque Sauce = "barbeque"
)

// Size : represents an available pizza size
type Size string

const (
	szSmall  Size = "small"
	szMedium Size = "medium"
	szLarge  Size = "large"
	szParty  Size = "party"
)

// Status : represents the status of an order
type Status string

const (
	stRecieved Status = "Order Recieved!"
	stAssembly Status = "Your pizza is being assembled!"
	stBaking   Status = "Your pizza is baking!"
	stReady    Status = "Your pizza is ready for pickup!"
	stEnRoute  Status = "Your pizza is on the way!"
)

// Topping : represents an avialable pizza topping
type Topping string

const (
	tpAnchovies Topping = "anchovies"
	tpArtichoke Topping = "artichoke"
	tpBasil     Topping = "basil"
	tpChicken   Topping = "chicken"
	tpHam       Topping = "ham"
	tpKale      Topping = "kale"
	tpOlives    Topping = "olives"
	tpOnion     Topping = "onion"
	tpPepperoni Topping = "pepperoni"
	tpPineapple Topping = "pineapple"
	tpTomato    Topping = "tomato"
)
