package main

import (
	"log"
	"net/http"

	"github.com/PixR2/wine-o/gql"
	"github.com/pkg/errors"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	rtdb "gopkg.in/gorethink/gorethink.v4"
)

func main() {
	session, err := rtdb.Connect(rtdb.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to rethinkdb"))
	}

	// Schema
	mutationSchema := gql.GetMutationSchema(session)
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(*mutationSchema)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	err = http.ListenAndServe("localhost:9078", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to listen for incomming connections"))
	}
}
