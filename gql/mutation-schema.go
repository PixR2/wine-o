package gql

import (
	"github.com/PixR2/wine-o/model"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	r "gopkg.in/gorethink/gorethink.v4"
)

func GetMutationSchema(db *r.Session) *graphql.ObjectConfig {
	var mutationSchema = &graphql.ObjectConfig{
		Name:        "wineo_mutations",
		Description: "Root of the WineO Mutation Schema",
		Fields: graphql.Fields{
			"createAccount": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"first_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"last_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// 1. Create account in database, setting it to unverified.
					err := model.CreateAccount(&model.Account{
						Email:     params.Args["email"].(string),
						FirstName: params.Args["first_name"].(string),
						LastName:  params.Args["last_name"].(string),
					}, db)
					if err != nil {
						return false, errors.Wrap(err, "createAccount mutation failed")
					}
					// 2. Send verification mail to accoutn email address.
					// email, _ := params.Args["email"].(string)

					return true, nil
				},
			},
		},
	}

	return mutationSchema
}
