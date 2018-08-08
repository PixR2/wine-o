package gql

import "github.com/graphql-go/graphql"

var WineObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Wine",
	Description: "Represents a wine of a certain year and winery in WineO.",
	Fields: graphql.Fields{
		"winery": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},

		"vintage": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"peak": &graphql.Field{
			Type: graphql.Int,
		},
		"hold": &graphql.Field{
			Type: graphql.Int,
		},
		"drink": &graphql.Field{
			Type: graphql.Int,
		},
		"taste": &graphql.Field{
			Type: graphql.Int,
		},

		"bottle_size": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"price": &graphql.Field{
			Type: graphql.Int,
		},

		"region": &graphql.Field{
			Type: graphql.String,
		},
		"varietal": &graphql.Field{
			Type: graphql.String,
		},
	},
})
