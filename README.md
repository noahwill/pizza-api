# ZaRnnr: pizza-api 
*A small pizza shop API built entirely in Go that sits on AWS DynamoDB and is deployed through Heroku*

## Project Structure
```
main.go
go.mod
go.sum
├── internal 
|    ├── valid
|    |   ├── accounts.go -- account create/update input validation
|    |   ├── misc.go     -- misc validation
|    |   └── orders.go   -- order create/update input validation
|    └── routes 
|        └── v1 -- Get all / Get one / Create / Update / Delete accounts/orders routes
|            ├── accounts.go 
|            └── orders.go  
├── pkg 
|    ├── client -- functions that facilitate calls to the accounts/orders routes
|    |   ├── accounts.go  
|    |   └── orders.go    
|    ├── helpers
|    |    └── http-helper.go -- facilitates an http request to the api for the client
|    └── types
|        ├── accounts.go -- account structure; account routes input/output types
|        ├── orders.go   -- order structure; order routes input/output types
|        └── subtypes.go -- misc types that are stored on accounts/orders
├── utils 
|   ├── db.go    -- database configuration
|   └── utils.go -- heartbeat route
└── vendor -- vendored dependencies
```

## Internal vs Pkg
When desiging this api, I saw the possibility for two different types of clients: one built in Go and one built in any other language. To make life easy for someone working in Go to interact with this api, I made the pkg directory that they could simply import into their code and use to easily make calls to the api endpoints. For any other client, if they don't want to shell out to running a small Go app that uses pkg, they are free to hit the endpoints implemented in internal directly. 

## Accounts + Orders
There are two underlying objects that back this api: accounts and orders. Both objects have Address objects.
```
type Account struct {             type Order struct {
	Active      bool                  Active      bool
	Address     Address               AccountID   string   
	CreatedAt   int64   	          Address     Address
	Email       string                CreatedAt   int64    
	FirstName   string                Price       float64 
	LastName    string                Delivery    bool
	LastUpdated int64                 LastUpdated int64
	Orders      []string              Size        string
	Password    string                Status      Status  
	UUID        string                Toppings    Toppings
}                                         UUID        string
                                  }

type Address struct {             type Toppings struct {
	ExtendedAddress string            Cheese   string   
	Locality        string            Sauce    string   
	PostalCode      string            Toppings []string 
	Region          string    }
	StreetAddress   string 
}
```

### Routes
There are five routes for each object stucture: Get (all), Get (one), Create, Update, and Delete. Below are details on the routes and some examples of `curl` commands to test them. (I'm running on Windows, my Mac/Linux curls are untested, let me know if you find any mistakes!)

#### Account Routes
- GET (Get all): takes no input at all; will return all active accounts

> Mac/Linux/Windows: `curl -X GET http://zarnnr.herokuapp.com/api/v1/account`
  
- GET (Get one): finds an account UUID in the url and will return one account if found for the id; will error if no account matches that id

> Mac/Linux/Windows: `curl -X GET http://zarnnr.herokuapp.com/api/v1/account/:uuid` *(I'd suggest getting all the accounts to find an id to use as an arg here)*
  
- POST (Create): looks for json input containing an `Address` object, an `Email` string, `FirstName` and `LastName` strings, and a `Password` string; all the fields are required
  - `Address`
    - `StreetAddress` must be alphanumeric and may contain spaces
    - `ExtendedAddress` may must alphanumeric, may contain spaces, and may also contain an octothorpe (#)
    - `Locality` and `Region` must be alphabetic and may contain spaces
    - `PostalCode` must only be numeric
  - `Email` must contain an @ symbol, have a valid format, be from a valid host, and exist at that host
  - `FirstName` and `LastName` must be alphabetic and may contain spaces
  - `Password` must be alphanumeric and may contain spaces (I'm a noob when it comes to security - that's a growth point from this project for sure)

> Mac/Linux: `curl -X POST -H "Content-Type: application/json" -d '{"Address": {"StreetAddress": "90 N Austin Ave", "ExtendedAddress": "#1E", "Locality": "Chicago", "Region": "IL", "PostalCode": "60607"}, "Email": "[YOUR EMAIL]", "FirstName": "Popa", "LastName": "John", "Password": "pizzatime" }' http://zarnnr.herokuapp.com/api/v1/account`

> Windows: `curl -X POST -H "Content-Type: application/json" -d "{\"Address\": {\"StreetAddress\": \"90 N Austin Ave\", \"ExtendedAddress\": \"#1E\", \"Locality\": \"Chicago\", \"Region\": \"IL\", \"PostalCode\": \"60607\"}, \"Email\": \"[YOUR EMAIL]\", \"FirstName\": \"Popa\", \"LastName\": \"John\", \"Password\": \"pizzatime\" }" http://zarnnr.herokuapp.com/api/v1/account`

- PUT (Update): finds an account UUID in the url and looks for json input that may contain a boolean `Active`, an updated `Address`, `Email`, `FirstName` and `LastName`, or `Password`; all the fields are optional; when updating an address, all Address fields must be populated (see above for validation details for these inputs); will error if no match is found for the given account id

> Mac/Linux: `curl -X PUT -H "Content-Type: application/json" -d '{"Active": false, "Address": {"StreetAddress": "UPDATE", "ExtendedAddress": "UPDATE", "Locality": "UPDATE", "Region": "UPDATE", "PostalCode": "12345"}, "Email": "[YOUR NEW EMAIL]", "FirstName": "UPDATE", "LastName": "UPDATE", "Password": "UPDATE" }' http://zarnnr.herokuapp.com/api/v1/account/:uuid`

> Windows: `curl -X PUT -H "Content-Type: application/json" -d "{\"Active\": false, \"Address\": {\"StreetAddress\": \"UPDATE\", \"ExtendedAddress\": \"UPDATE\", \"Locality\": \"UPDATE\", \"Region\": \"UPDATE\", \"PostalCode\": \"12345\"}, \"Email\": \"[YOUR NEW EMAIL]\", \"FirstName\": \"UPDATE\", \"LastName\": \"UPDATE\", \"Password\": \"UPDATE\" }" http://zarnnr.herokuapp.com/api/v1/account/:uuid`

- DELETE: finds an account UUID in the url and will delete one account if found and deletes any orders made by that account

> Mac/Linux/Windows: `curl -X DELETE http://zarnnr.herokuapp.com/api/v1/account/:uuid`

#### Order Routes

- GET (Get all): finds an account UUID in the url; json object with an `Active` field is optional; true will return all active orders, false will return inactive orders, nil (unspecified) will return all orders for the account; it will error with an invalid account UUID

> Mac/Linux: `curl -X GET -H "Content-Type: application/json" -d '{ "Active": true }' http://zarnnr.herokuapp.com/api/v1/:account/order`

> Windows: `curl -X GET -H "Content-Type: application/json" -d "{ \"Active\": true }" http://zarnnr.herokuapp.com/api/v1/:account/order`

- GET (Get one): finds for an account and order UUID in the url; will return the order if found, otherwise it will error

> Mac/Linux/Windows: `curl -X GET http://zarnnr.herokuapp.com/api/v1/:account/order/:uuid`

- POST (Create): looks for json input containing an `Address` object, an `Delivery` boolean, an order `Size` (string), and a `Toppings` object; `Deliver`, `Size`, and `Toppings` are required
  - `Address` is validated in the same way as laid out above and is not required. `Delivery` is required, when true and `Address` is nil, the order will be assigned the account's default address. When `Delivery` is false but `Address` is specified, the order address will not be assigned.
  - `Size` has 4 valid options: small, medium, large, and party ($8, $16, $32, and $48 respectively)
  - `Toppings`
    - `Cheese` has 3 valid options: cheddar, mozzarella, and parmesean ($3, $3, and $4)
    - `Sauce` has 3 valid options: tomato, white, and barbeque ($2, $2, and $3)
    - `Toppings` can be of any length and has 11 valid options: anchovies ($0.75), artichokes ($1.25), basil ($0.75), chicken ($1.75), ham ($1.25), kale ($0.75), olives ($0.50), onion ($0.50), pepperoni ($1.25), pineapple ($100), and tomato ($0.75)
   - After validation, the order will sum all of the prices together to populate it's `Price` field
   
> Mac/Linux: `curl -X POST -H "Content-Type: application/json" -d '{ "Address": { "StreetAddress": "STREET", "ExtendedAddress": "EXTENDED", "Locality": "LOCALITY", "Region": "REGION", "PostalCode": "12345" }, "Delivery": true, "Size": "party", "Toppings": { "Cheese": "parmesan", "Sauce": "barbeque", "Toppings": ["ham", "kale", "pineapple"] } }' http://zarnnr.herokuapp.com/api/v1/:account/order`

> Windows: `curl -X POST -H "Content-Type: application/json" -d "{ \"Address\": { \"StreetAddress\": \"STREET\", \"ExtendedAddress\": \"EXTENDED\", \"Locality\": \"LOCALITY\", \"Region\": \"REGION\", \"PostalCode\": \"12345\" }, \"Delivery\": true, \"Size\": \"party\", \"Toppings\": { \"Cheese\": \"parmesan\", \"Sauce\": \"barbeque\", \"Toppings\": [\"ham\", \"kale\", \"pineapple\"] } }" http://zarnnr.herokuapp.com/api/v1/:account/order`

- PUT (Update): finds an account UUID and order UUID in the url and looks for json input that may contain an `Active` boolean, an updated `Address` and `Delivery` boolean, an order `Size` (string), an order `Status` (string), or a `Toppings` object; all the fields are optional; when updating an address, all `Address` fields must be populated and `Delivery` must be true (see above for validation details for these inputs); validation for the other fields is as above; will error if no match is found for the given account id
  - `Status` has 7 valid inputs: Received, Assembly, Baking, Ready, PickedUp, EnRoute, and Delivered. This is more of an internal maintenance datum, it would be the pizza shop updating the `Status` as the pizza is being made. If the current Order status is not Received, then `Size` and `Toppings` cannot be updated. If the current Order status is PickedUp, EnRoute, or Delivered, then `Delivery` and `Address` cannot be updated
  - Order `Price` will be updated according to new inputs
  
> Mac/Linux: `curl -X PUT -H "Content-Type: application/json" -d '{ "Active": false, "Address": { "StreetAddress": "UPDATE", "ExtendedAddress": "UPDATE", "Locality": "UPDATE", "Region": "UPDATE", "PostalCode": "54321" }, "Delivery": true, "Size": "large", "Status": "Received", "Toppings": { "Cheese": "cheddar", "Sauce": "tomato", "Toppings": ["anchovies", "artichoke", "basil"] } }' http://zarnnr.herokuapp.com/api/v1/:account/order/:uuid`

> Windows: `curl -X PUT -H "Content-Type: application/json" -d "{ \"Active\": false, \"Address\": { \"StreetAddress\": \"UPDATE\", \"ExtendedAddress\": \"UPDATE\", \"Locality\": \"UPDATE\", \"Region\": \"UPDATE\", \"PostalCode\": \"54321\" }, \"Delivery\": true, \"Size\": \"large\", \"Status\": \"Received\", \"Toppings\": { \"Cheese\": \"cheddar\", \"Sauce\": \"tomato\", \"Toppings\": [\"anchovies\", \"artichoke\", \"basil\"] } }" http://zarnnr.herokuapp.com/api/v1/:account/order/:uuid`

- DELETE: finds an account and order UUID in the url and will delete one order and remove it from the list of order ids on the account object; will error if no order is found

> Mac/Linux/Windows: `curl -X DELETE http://zarnnr.herokuapp.com/api/v1/:account/order/:uuid`






















