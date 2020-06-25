package main

import (
	"os"
	v1routes "pizza-api/internal/routes/v1"
	"pizza-api/utils"
	"time"

	"github.com/apex/log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
)

func main() {
	utils.Config.SetConfigs()

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
	v1.GET("/:account/order", v1routes.GetAccountOrdersRoute)
	v1.GET("/:account/order/:uuid", v1routes.GetAccountOrderRoute)
	v1.POST("/:account/order", v1routes.CreateAccountOrderRoute)
	v1.PUT("/:account/order/:uuid", v1routes.UpdateAccountOrderRoute)
	v1.DELETE("/:account/order/:uuid", v1routes.DeleteAccountOrderRoute)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
