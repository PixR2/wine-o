package gql

import "github.com/graphql-go/graphql"

var WinePositionObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Wine Position",
	Description: "Represents a position of a certain wine in a collection of wines in WineO.",
	Fields: graphql.Fields{
		"collection_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"compartment": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"wine": &graphql.Field{
			Type: WineObject,
		},
	},
})
