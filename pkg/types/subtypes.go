package types

// Address : represents a ZaRnnr order address
type Address struct {
	StreetAddress   string `json:"startTime,omitempty"`
	ExtendedAddress string `json:"extendedAddress,omitempty"`
	CreatedAt       int64  `json:"createdAt,omitempty"`
	LastUpdated     int64  `json:"lastUpdated,omitempty"`
	Locality        string `json:"locality,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	Region          string `json:"region,omitempty"`
	Country         string `json:"country,omitempty"`
	UUID            string `json:"UUID,omitempty"`
}

// Toppings : represents a ZaRnnr order toppings
type Toppings struct {
	Cheese   Cheese    `json:"cheese,omitempty"`
	Sauce    Sauce     `json:"sauce,omitempty"`
	Toppings []Topping `json:"toppings,omitempty"`
}

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

// Cheese : represents an available pizza cheese
type Cheese string

const (
	chCheddar    Cheese = "cheddar"
	chMozzarella Cheese = "mozzarella"
	chParmesan   Cheese = "parmesan"
)

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

type Status string

const (
	stRecieved Status = "Order Recieved!"
	stAssembly Status = "Your pizza is being assembled!"
	stBaking   Status = "Your pizza is baking!"
	stReady    Status = "Your pizza is ready for pickup!"
	stEnRoute  Status = "Your pizza is on the way!"
)
