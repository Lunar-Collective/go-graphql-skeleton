package store

import (
	"context"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)


func ArangoInit(ctx context.Context, host, dbName string) (driver.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{host},
	})

	if err != nil {
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})

	if err != nil {
		return nil, err
	}

	db, err := c.Database(ctx, dbName)

	return db, err
}