package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	wineo_api_service "github.com/pixr2/wineo/wineo-api-service"
	rethink "gopkg.in/gorethink/gorethink.v4"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	session, err := rethink.Connect(rethink.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "wineo",
	})
	if err != nil {
		panic(err)
		return
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(wineo_api_service.NewExecutableSchema(wineo_api_service.Config{Resolvers: wineo_api_service.NewResolver(session)})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
