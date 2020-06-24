package main

import (
	"database/sql"
	"fmt"
	"os"
	v1routes "pizza-api/internal/routes/v1"
	"pizza-api/utils"
	"time"

	"github.com/apex/log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
)

// Builds connection to the orders and accounts dbs
// (init are auto called by Go)
func init() {
	uri := "%s:%s@tcp(%s:%s)/%s?sslMode=disable"
	accountsURI := "vxbeicwalhzehb:8648927891653f4cf76a2c664256971723b1ef98569acc5a680843ebf02ffe1d@ec2-34-192-173-173.compute-1.amazonaws.com:5432/d8m251cskkpkf8"
	// fmt.Sprintf(uri, utils.Username, utils.AccountsPassword, utils.AccountsHost, utils.Port, utils.AccountsDB)
	addressesURI := fmt.Sprintf(uri, utils.Username, utils.AddressesPassword, utils.AddressesHost, utils.Port, utils.AddressesDB)
	ordersURI := fmt.Sprintf(uri, utils.Username, utils.OrdersPassword, utils.OrdersHost, utils.Port, utils.OrdersDB)
	toppingsURI := fmt.Sprintf(uri, utils.Username, utils.ToppingsPassword, utils.ToppingsHost, utils.Port, utils.ToppingsDB)

	log.Infof("| Accounts  URI: %s", accountsURI)
	log.Infof("| Addresses URI: %s", addressesURI)
	log.Infof("| Orders    URI: %s", ordersURI)
	log.Infof("| Toppings  URI: %s", toppingsURI)

	accountsConn, err := sql.Open("postgres", accountsURI)
	if err != nil {
		log.Errorf("| Could not open connection to the accounts db with error: %s", err.Error())
	} else if err := accountsConn.Ping(); err != nil {
		log.Errorf("| Could not open connection to the accounts db with error: %s", err.Error())
	} else {
		log.Infof("| Successfully connected to the accounts db!")
	}

	addressesConn, err := sql.Open("postgres", addressesURI)
	if err != nil {
		log.Errorf("| Could not open connection to the addresses db with error: %s", err.Error())
	} else if err := addressesConn.Ping(); err != nil {
		log.Errorf("| Could not open connection to the addresses db with error: %s", err.Error())
	} else {
		log.Infof("| Successfully connected to the addresses db!")
	}

	ordersConn, err := sql.Open("postgres", ordersURI)
	if err != nil {
		log.Errorf("| Could not open connection to the orders db with error: %s", err.Error())
	} else if err := ordersConn.Ping(); err != nil {
		log.Errorf("| Could not open connection to the orders db with error: %s", err.Error())
	} else {
		log.Infof("| Successfully connected to the orders db!")
	}

	toppingsConn, err := sql.Open("postgres", toppingsURI)
	if err != nil {
		log.Errorf("| Could not open connection to the toppings db with error: %s", err.Error())
	} else if err := toppingsConn.Ping(); err != nil {
		log.Errorf("| Could not open connection to the toppings db with error: %s", err.Error())
	} else {
		log.Infof("| Successfully connected to the toppings db!")
	}

	utils.SetAccountsDB(accountsConn)
	utils.SetAddressesDB(addressesConn)
	utils.SetOrdersDB(ordersConn)
	utils.SetToppingsDB(toppingsConn)

	// Ensure that the structure of the tables matches the structure of the types they support
	// utils.AssertAccountsSchema()
	// utils.AssertAddressesSchema()
	// utils.AssertOrdersSchema()
	// utils.AssertToppingsSchema()
}

func main() {
	log.Infof("| ZaRnnr is hot 'n ready!")

	e := echo.New()
	e.Server.IdleTimeout = 30 * time.Second
	e.Server.ReadTimeout = 15 * time.Second
	e.Server.ReadHeaderTimeout = 10 * time.Second

	// utils
	e.GET("/api/v1/heartbeat", utils.HeartbeatRoute)

	// V1 API Routes
	// if we want to make changes to routes / make new routes to test in production
	// we can create a v2 so that the downstream consumers of this api can have uninterupted
	// service and have ample time to switch over to the new routes after a release
	v1 := e.Group("/api/v1")

	// accounts
	v1.GET("/account", v1routes.GetAccountsRoute)
	v1.GET("/account/:uuid", v1routes.GetAccountRoute)
	v1.POST("/account", v1routes.CreateAccountRoute)
	v1.PUT("/account/:uuid", v1routes.UpdateAccountRoute)
	v1.DELETE("/account/:uuid", v1routes.DeleteAccountRoute)

	// orders
	v1.GET("/order", v1routes.GetOrdersRoute)
	v1.GET("/order/:uuid", v1routes.GetOrderRoute)
	v1.POST("/order", v1routes.CreateOrderRoute)
	v1.PUT("/order/:uuid", v1routes.UpdateOrderRoute)
	v1.DELETE("/order/:uuid", v1routes.DeleteOrderRoute)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
