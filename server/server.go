package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	golang_graphql_user_mgr "github.com/BolajiOlajide/golang-graphql-user-mgr"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	golang_graphql_user_mgr.NewDatabase()

	router := chi.NewRouter()
	router.Use(golang_graphql_user_mgr.Middleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(golang_graphql_user_mgr.NewExecutableSchema(golang_graphql_user_mgr.Config{Resolvers: &golang_graphql_user_mgr.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
