package utils

import (
	"github.com/jinzhu/gorm"
)

var accounts, orders, toppings *gorm.DB

const (
	// Fill in local postgres username and password and create accounts, orders, and toppings dbs before running

	// Username : local postgres username
	Username = "postgres"

	// Password : local postgres password
	Password = "N0@hmaryc0rd"

	// AccountsDB : name of local postgres accounts db
	AccountsDB = "accounts"

	// OrdersDB : name of local postgres orders db
	OrdersDB = "orders"

	// ToppingsDB : name of local postgress orders db
	ToppingsDB = "toppings"

	// DBHost : dbhost
	DBHost = "localhost"
)

// SetAccountsDB : sets the accounts database
func SetAccountsDB(a *gorm.DB) { accounts = a }

// GetAccountsDB : returns the accounts database
func GetAccountsDB() *gorm.DB { return accounts }

// AssertAccountsSchema : ensures that the accounts db is using the correct schema
func AssertAccountsSchema() { accounts.Debug().AutoMigrate(&Account{}) }

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
