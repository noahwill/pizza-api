package client

// Toppings : represents the toppings on a pizza
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
