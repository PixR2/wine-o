package gql

import "github.com/graphql-go/graphql"

var WineCollectionObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Wine Collection",
	Description: "Represents a collection of wine positions in WineO.",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"owners": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(AccountObject)),
		},
	},
})
