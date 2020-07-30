package main

import (
	"context"
	"github.com/Lunar-Collective/go-graphql-skeleton/store"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Lunar-Collective/go-graphql-skeleton/graph"
	"github.com/Lunar-Collective/go-graphql-skeleton/graph/generated"
)

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv(("DB_HOST"))
	
	if port == "" {
		port = defaultPort
	}

	// Add new background context
	ctx := context.Background()

	db, err := store.ArangoInit(ctx, dbHost, dbName)

	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ArangoDB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
