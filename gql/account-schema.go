package gql

import "github.com/graphql-go/graphql"

var AccountObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Account",
	Description: "An account registered with WineO.",
	Fields: graphql.Fields{
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
