package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/makrorof/gorm-graphql-postgres-echo-auth/api/gql"
	"github.com/makrorof/gorm-graphql-postgres-echo-auth/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(cfg *configs.Config) {
	db, err := initDb(cfg)

	if err != nil {
		log.Panicln("DB not connected, Err:", err)
		return
	}

	_ = db

	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Application name: ", os.Getenv("APPLICATION_NAME"))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.APIConf.Port)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(cfg.APIConf.Port), nil))
}

func initDb(cfg *configs.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", cfg.PostgresConf.Username, cfg.PostgresConf.Password, cfg.PostgresConf.Name, cfg.PostgresConf.Host, cfg.PostgresConf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
