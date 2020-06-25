package utils

// import (
// 	"database/sql"
// )

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// Config for the app
var Config Configuration

// Configuration representation of the app
type Configuration struct {
	AccountsTableConn dynamo.Table
	OrdersTableConn   dynamo.Table
	DynamoSession     *dynamo.DB
}

// SetConfigs sets the configuration settings for the environment
func (c *Configuration) SetConfigs() {
	db := DynamoSessionInit()

	accountsTable := db.Table(os.Getenv("AccountsTable"))
	ordersTable := db.Table(os.Getenv("OrdersTable"))

	(*c).AccountsTableConn = accountsTable
	(*c).OrdersTableConn = ordersTable
	(*c).DynamoSession = db
}

// DynamoSessionInit open a session with dynamodb
func DynamoSessionInit() *dynamo.DB {
	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}
	sess := session.Must(session.NewSession(config))
	db := dynamo.New(sess, config)
	return db
}
