package utils

import (
	"github.com/jinzhu/gorm"
)

var accounts, addresses, orders, toppings *gorm.DB

const (
	// Fill in local postgres username and password and create accounts, orders, and toppings dbs before running

	// Username : local postgres username
	Username = "postgres"
	// Port : dbs port
	Port = "5432"

	// AccountsDB : name of accounts db
	AccountsDB = "accounts"
	// AccountsPassword : password of accounts db
	AccountsPassword = "sMTeKeqnb7ic8e4gQejD"
	// AccountsHost : accounts db host
	AccountsHost = "accounts.cx7blgdanz85.us-east-1.rds.amazonaws.com"

	// AddressesDB : name of addresses db
	AddressesDB = "addresses"
	// AddressesPassword : password of addresses db
	AddressesPassword = "WV5Qqv7hhuHuaRKu7MG6"
	// AddressesHost : addresses db host
	AddressesHost = "addresses.cx7blgdanz85.us-east-1.rds.amazonaws.com"

	// OrdersDB : name of orders db
	OrdersDB = "orders"
	// OrdersPassword : password of orders db
	OrdersPassword = "mPWm9xEtjV8KzhIFm6GI"
	// OrdersHost : orders db host
	OrdersHost = "orders.cx7blgdanz85.us-east-1.rds.amazonaws.com"

	// ToppingsDB : name of toppings db
	ToppingsDB = "toppings"
	// ToppingsPassword : password of toppings db
	ToppingsPassword = "AaIJVleEYv4J6QMAs7YT"
	// ToppingsHost : addresses db host
	ToppingsHost = "toppings.cx7blgdanz85.us-east-1.rds.amazonaws.com"
)

// SetAccountsDB : sets the accounts database
func SetAccountsDB(a *gorm.DB) { accounts = a }

// GetAccountsDB : returns the accounts database
func GetAccountsDB() *gorm.DB { return accounts }

// AssertAccountsSchema : ensures that the accounts db is using the correct schema
func AssertAccountsSchema() { accounts.Debug().AutoMigrate(&Account{}) }

// SetAddressesDB : sets the addresses database
func SetAddressesDB(a *gorm.DB) { addresses = a }

// GetAddressesDB : returns the addresses database
func GetAddressesDB() *gorm.DB { return addresses }

// AssertAddressesSchema : ensures that the addresses db is using the correct schema
func AssertAddressesSchema() { accounts.Debug().AutoMigrate(&Address{}) }

// SetOrdersDB : sets the orders database
func SetOrdersDB(o *gorm.DB) { orders = o }

// GetOrdersDB : returns the orders database
func GetOrdersDB() *gorm.DB { return orders }

// AssertOrdersSchema : ensures that the orders db is using the correct schema
func AssertOrdersSchema() { orders.Debug().AutoMigrate(&Order{}) }

// SetToppingsDB : sets the toppings database
func SetToppingsDB(t *gorm.DB) { toppings = t }

// GetToppingsDB : returns the accounts database
func GetToppingsDB() *gorm.DB { return toppings }

// AssertToppingsSchema : ensures that the toppings db is using the correct schema
func AssertToppingsSchema() { toppings.Debug().AutoMigrate(&Toppings{}) }
