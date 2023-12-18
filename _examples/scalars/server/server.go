package main

import (
	"log"
	"net/http"

	"github.com/ervitis/gqlgen/example/scalars"
	"github.com/ervitis/gqlgen/graphql/handler"
	"github.com/ervitis/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
