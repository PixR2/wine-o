package gql

import (
	"github.com/graphql-go/graphql"
	r "gopkg.in/gorethink/gorethink.v4"
)

func GetQuerySchema(db *r.Session) *graphql.Object {
	var querySchema = graphql.NewObject(graphql.ObjectConfig{
		Name:        "WineO GraphQL Query Schema",
		Description: "Root of the WineO Query Schema",
		Fields: graphql.Fields{
			"account": &graphql.Field{
				Type: AccountObject,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
			"wine_collections": &graphql.Field{
				Type: graphql.NewList(WineCollectionObject),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
			"wine": &graphql.Field{
				Type: WineObject,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
			"wine_positions": &graphql.Field{
				Type: graphql.NewList(WinePositionObject),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
		},
	})

	return querySchema
}
