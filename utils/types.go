package utils

// Account : represents a ZaRnnr account
type Account struct {
	Active      bool   `json:"active,omitempty"`
	CreatedAt   int64  `json:"createdAt,omitempty"`
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	LastUpdated int64  `json:"lastUpdated,omitempty"`
	Password    string `json:"password,omitempty"` // TODO: this pizza shop dev hasn't done much work with security
	UUID        string `json:"UUID,omitempty"`
}

// Order : represents a ZaRnnr order
type Order struct {
	AccountID   string `json:"accountID,omitempty"`
	AddressID   string `json:"addressID,omitempty"`
	Active      bool   `json:"active,omitempty"`
	Cheese      Cheese `json:"cheese,omitempty"`
	CreatedAt   int64  `json:"createdAt,omitempty"`
	LastUpdated int64  `json:"lastUpdated,omitempty"`
	Pickup      bool   `json:"pickup,omitempty"`
	Sauce       Sauce  `json:"sauce,omitempty"`
	Size        Size   `json:"size,omitempty"`
	ToppingsID  string `json:"toppingsID,omitempty"`
	UUID        string `json:"UUID,omitempty"`
}

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
	Anchovies  bool
	Artichokes bool
	Basil      bool
	Chicken    bool
	Ham        bool
	Kale       bool
	Olives     bool
	Onion      bool
	Pepperoni  bool
	Pineapple  bool
	Tomato     bool
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
