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

