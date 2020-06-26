# ZaRnnr: pizza-api
*A small pizza shop API built entirely in Go that sits on AWS DynamoDB and is deployed through Heroku*

## Project Structure
```
main.go
├── internal 
|    ├── validate
|    |   ├── accounts.go // account create/update input validation
|    |   ├── misc.go     // misc validation
|    |   └── orders.go   // order create/update input validation
|    └── routes 
|        └── v1 // Get (all) / Get (one) / Create / Update / Delete accounts/orders routes
|            ├── accounts.go 
|            └── orders.go  
├── pkg 
|    ├── client // functions that facilitate calls to the accounts/orders routes
|    |   ├── accounts.go  
|    |   └── orders.go    
|    ├── helpers
|    |    ├── http-helper.go // Facilitates an http request for the client
|    └── types
|        ├── accounts.go // account structure; account routes input/output types
|        ├── orders.go   // order structure; order routes input/output types
|        └── subtypes.go // misc types that are stored on accounts/orders
├── utils 
|   ├── db.go    // database configuration
|   └── utils.go // heartbeat route
└── vendor // vendored dependencies
```

## Internal vs Pkg
When desiging this api, I saw the possibility for two different types of clients: one built in Go and one built in any other language. To make life easy for someone working in Go to interact with this api, I made the pkg directory that they could simply import into their code and use to easily make calls to the api endpoints. For any other client, if they don't want to shell out to running a small Go app that uses pkg, they are free to hit the endpoints implemented in internal directly. 

## Accounts + Orders
There are two underlying objects that back this api: accounts and orders. Both objects have Address objects.
```
type Account struct {
	Active      bool    
	Address     Address 
	CreatedAt   int64   
	Email       string  
	FirstName   string  
	LastName    string  
	LastUpdated int64   
	Orders      []string
	Password    string  
	UUID        string  
}

type Order struct {
	Active      bool     
	AccountID   string   
	Address     Address  
	CreatedAt   int64    
	Price       float64  
	Delivery    bool     
	LastUpdated int64    
	Size        string   
	Status      Status   
	Toppings    Toppings 
	UUID        string   
}

type Address struct {
	ExtendedAddress string 
	Locality        string 
	PostalCode      string 
	Region          string 
	StreetAddress   string 
}
```

## Routes
There are five routes for each object stucture: Get (all), Get (one), Create, Update, and Delete. Below are details on the routes and some examples of `curl` commands to test them.

### Accounts
- GET (Get all): a very simple route that takes no input at all; will return all active accounts

> Mac/Linux/Windows: `curl -X GET http://zarnnr.herokuapp.com/api/v1/account`
  
- GET (Get one): finds an account id in the url and will return one account if found for the id; will error if no account matches that id

> Mac/Linux/Windows: `curl -X GET http://zarnnr.herokuapp.com/api/v1/account/:uuid` *(I'd suggest getting all the accounts to find an id to use as an arg here)*
  
- POST (Create): looks for a json input containing an Address object, an email string, first and last name strings, and a password string; all the fields are required
  - `Address`
    - `StreetAddress` must be alphanumeric and may contain spaces
    - `ExtendedAddress` may must alphanumeric, may contain spaces, and may also contain an octothorpe (#)
    - `Locality` and `Region` must be alphabetic and may contain spaces
    - `PostalCode` must only be numeric
  - `Email` must contain an @ symbol, exist, have a valid format, and be from a valid host
  - `FirstName` and `LastName` must be alphabetic and may contain spaces
  - `Password` must be alphanumeric and may contain spaces (I'm a noob when it comes to security - that's a growth point from this project for sure)

> Mac/Linux: `curl -X POST -H "Content-Type: application/json" -d '{"Address": {"StreetAddress": "90 N Austin Ave", "ExtendedAddress": "#1E", "Locality": "Chicago", "Region": "IL", "PostalCode": "60607"}, "Email": "[YOUR EMAIL]", "FirstName": "Popa", "LastName": "John", "Password": "pizzatime" }' http://zarnnr.herokuapp.com/api/v1/account`

> Windows: `curl -X POST -H "Content-Type: application/json" -d "{\"Address\": {\"StreetAddress\": \"90 N Austin Ave\", \"ExtendedAddress\": \"#1E\", \"Locality\": \"Chicago\", \"Region\": \"IL\", \"PostalCode\": \"60607\"}, \"Email\": \"[YOUR EMAIL]\", \"FirstName\": \"Popa\", \"LastName\": \"John\", \"Password\": \"pizzatime\" }" http://zarnnr.herokuapp.com/api/v1/account`

- PUT (Update): finds an account id in the url and looks for a json input that may contain a boolean Active, an updated Address, email, first and last name, or password; all the fields are optional; when updating an address, all Address fields must be populated (see above for validation details for these inputs); will error if no match is found for the given account id

> Mac/Linux: `curl -X PUT -H "Content-Type: application/json" -d '{"Active": false, "Address": {"StreetAddress": "UPDATE", "ExtendedAddress": "UPDATE", "Locality": "UPDATE", "Region": "UPDATE", "PostalCode": "12345"}, "Email": "[YOUR NEW EMAIL]", "FirstName": "UPDATE", "LastName": "UPDATE", "Password": "UPDATE" }' http://zarnnr.herokuapp.com/api/v1/account/:uuid`

> Windows: `curl -X PUT -H "Content-Type: application/json" -d "{\"Active\": false, \"Address\": {\"StreetAddress\": \"UPDATE\", \"ExtendedAddress\": \"UPDATE\", \"Locality\": \"UPDATE\", \"Region\": \"UPDATE\", \"PostalCode\": \"12345\"}, \"Email\": \"[YOUR NEW EMAIL]\", \"FirstName\": \"UPDATE\", \"LastName\": \"UPDATE\", \"Password\": \"UPDATE\" }" http://zarnnr.herokuapp.com/api/v1/account/c54e54df-0812-424e-ade6-b28ac9a8cf48`

- DELETE: finds an account id in the url and will delete one account if found and deletes any orders made by that account

> Mac/Linux/Windows: `curl -X DELETE http://zarnnr.herokuapp.com/api/v1/account/:uuid`
