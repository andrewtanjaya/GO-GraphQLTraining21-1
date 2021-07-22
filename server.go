package main

import (
	"Go-Graphql/graph"
	"Go-Graphql/graph/generated"
	"Go-Graphql/postgre"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db := postgre.New(&pg.Options{
		Addr:     ":5440",
		User:     "postgres",
		Password: "root",
		Database: "SOolaria",
	})
	db.AddQueryHook(postgre.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
log.Fatal(http.ListenAndServe(":"+port, router))
}
