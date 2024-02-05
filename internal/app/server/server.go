package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/makrorof/gorm-graphql-postgres-echo-auth/api/gql"
)

const defaultPort = "8080"

func Run() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Application name: ", os.Getenv("APPLICATION_NAME"))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
