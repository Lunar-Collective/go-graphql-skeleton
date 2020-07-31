package graph

import "github.com/arangodb/go-driver"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	ArangoDB driver.Database
	CheckScope func(scope, authDomain, token string) bool
}
