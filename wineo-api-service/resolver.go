//go:generate gorunpkg github.com/99designs/gqlgen

package wineo_api_service

import (
	context "context"

	"github.com/pkg/errors"
	rethink "gopkg.in/gorethink/gorethink.v4"
)

type Resolver struct {
	rtdb *rethink.Session
}

func NewResolver(session *rethink.Session) *Resolver {
	return &Resolver{
		rtdb: session,
	}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddWine(ctx context.Context, winery string, varietal string, vintage int, bottleSize int) (*Wine, error) {
	_, err := rethink.Table("wines").Insert(&Wine{
		Winery:     winery,
		Varietal:   varietal,
		Vintage:    vintage,
		BottleSize: bottleSize,
	}).RunWrite(r.rtdb)

	if err != nil {
		return nil, errors.Wrap(err, "failed to insert new wine into database")
	}

	return nil, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Wine(ctx context.Context, winery string, vintage int, bottleSize int) (*Wine, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) WineAdded(ctx context.Context) (<-chan Wine, error) {
	c := make(chan Wine)
	res, err := rethink.Table("wines").Changes().Field("new_val").Run(r.rtdb)
	if err != nil {
		return nil, errors.Wrap(err, "failed to subscribe to change feed of wines table")
	}

	res.Listen(c)

	return c, nil
}

type wineResolver struct{ *Resolver }

func (r *wineResolver) BottleSize(ctx context.Context, obj *Wine) (int, error) {
	panic("not implemented")
}
