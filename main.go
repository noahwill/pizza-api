package main

import (
	"database/sql"
	"fmt"
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
	var uri = "%s:%s@tcp(%s:%s)/%s="
	accountsURI := fmt.Sprintf(uri, utils.Username, utils.AccountsPassword,
		utils.AccountsHost, utils.Port, utils.AccountsDB)
	// accountsURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	// utils.AccountsHost, utils.Port, utils.Username, utils.AccountsDB, utils.AccountsPassword)
	// addressesURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
	// 	utils.AddressesHost, utils.Port, utils.Username, utils.AddressesDB, utils.AddressesPassword)
	// ordersURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
	// 	utils.OrdersHost, utils.Port, utils.Username, utils.OrdersDB, utils.OrdersPassword)
	// toppingsURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
	// 	utils.ToppingsHost, utils.Port, utils.Username, utils.ToppingsDB, utils.ToppingsPassword)

	log.Infof("| Accounts  URI: %s", accountsURI)
	// log.Infof("| Addresses URI: %s", addressesURI)
	// log.Infof("| Orders    URI: %s", ordersURI)
	// log.Infof("| Toppings  URI: %s", toppingsURI)

	accountsConn, err := sql.Open("postgres", accountsURI)
	defer accountsConn.Close()
	if err != nil {
		log.Errorf("| Could not open connection to the accounts db with error: %s", err.Error())
	} else {
		log.Infof("| Successfully connected to the accounts db!")
	}

	// addressesConn, err := gorm.Open("postgres", addressesURI)
	// defer addressesConn.Close()
	// if err != nil {
	// 	log.Errorf("| Could not open connection to the addresses db with error: %s", err.Error())
	// } else {
	// 	log.Infof("| Successfully connected to the addresses db!")
	// }

	// ordersConn, err := gorm.Open("postgres", ordersURI)
	// defer ordersConn.Close()
	// if err != nil {
	// 	log.Errorf("| Could not open connection to the orders db with error: %s", err.Error())
	// } else {
	// 	log.Infof("| Successfully connected to the orders db!")
	// }

	// toppingsConn, err := gorm.Open("postgres", toppingsURI)
	// defer toppingsConn.Close()
	// if err != nil {
	// 	log.Errorf("| Could not open connection to the toppings db with error: %s", err.Error())
	// } else {
	// 	log.Infof("| Successfully connected to the toppings db!")
	// }

	utils.SetAccountsDB(accountsConn)
	// utils.SetAddressesDB(addressesConn)
	// utils.SetOrdersDB(ordersConn)
	// utils.SetToppingsDB(toppingsConn)

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

	e.Logger.Fatal(e.Start(":8282"))
}
