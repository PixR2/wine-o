//go:generate gorunpkg github.com/99designs/gqlgen

package wineo_api_service

import (
	context "context"
)

type Resolver struct{}

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

func (r *mutationResolver) Signup(ctx context.Context, email string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) VerifyEmail(ctx context.Context, token string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) Signin(ctx context.Context, email string, password string) (*string, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, email *string, firstName *string, lastName *string, password *string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) RemoveUser(ctx context.Context, deleteCollections bool, reason *string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateCollection(ctx context.Context, name string) (*string, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateCollection(ctx context.Context, id string, name string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) RemoveCollection(ctx context.Context, id string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddPosition(ctx context.Context, collectionID string, winery string, varietal string, vintage int, bottleSize int, numBottles int, compartment *string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePosition(ctx context.Context, collectionID string, winery string, varietal string, vintage int, bottleSize int, numBottles int, compartment *string) (*bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (User, error) {
	panic("not implemented")
}
func (r *queryResolver) Wine(ctx context.Context, winery string, varietal string, vintage int, bottleSize int) (*Wine, error) {
	panic("not implemented")
}
func (r *queryResolver) Collection(ctx context.Context, id string) (*WineCollection, error) {
	panic("not implemented")
}
func (r *queryResolver) MyCollections(ctx context.Context) ([]*WineCollection, error) {
	panic("not implemented")
}
func (r *queryResolver) Positions(ctx context.Context, collectionID string) ([]*WinePosition, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) MeChanged(ctx context.Context) (<-chan *User, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) PositionChanged(ctx context.Context, collectionID string) (<-chan *WinePosition, error) {
	panic("not implemented")
}
