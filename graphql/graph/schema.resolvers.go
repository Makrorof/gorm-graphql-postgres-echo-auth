package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/makrorof/gorm-graphql-postgres-echo-auth/graphql/graph/model"
)

// CreateRequest is the resolver for the createRequest field.
func (r *mutationResolver) CreateRequest(ctx context.Context, input model.ScanRequestInput) (*model.ScanRequest, error) {
	panic(fmt.Errorf("not implemented: CreateRequest - createRequest"))
}

// DeleteRequest is the resolver for the deleteRequest field.
func (r *mutationResolver) DeleteRequest(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteRequest - deleteRequest"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// CreatePlan is the resolver for the createPlan field.
func (r *mutationResolver) CreatePlan(ctx context.Context, input model.PlanInput) (*model.Plan, error) {
	panic(fmt.Errorf("not implemented: CreatePlan - createPlan"))
}

// UpdatePlan is the resolver for the updatePlan field.
func (r *mutationResolver) UpdatePlan(ctx context.Context, input model.PlanInput) (*model.Plan, error) {
	panic(fmt.Errorf("not implemented: UpdatePlan - updatePlan"))
}

// DeletePlan is the resolver for the deletePlan field.
func (r *mutationResolver) DeletePlan(ctx context.Context, planID string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeletePlan - deletePlan"))
}

// SelectPlan is the resolver for the selectPlan field.
func (r *mutationResolver) SelectPlan(ctx context.Context, userID string, planID string) (*model.UserPlan, error) {
	panic(fmt.Errorf("not implemented: SelectPlan - selectPlan"))
}

// CancelPlan is the resolver for the cancelPlan field.
func (r *mutationResolver) CancelPlan(ctx context.Context, userID string, planID string) (bool, error) {
	panic(fmt.Errorf("not implemented: CancelPlan - cancelPlan"))
}

// Requests is the resolver for the requests field.
func (r *queryResolver) Requests(ctx context.Context) ([]*model.ScanRequest, error) {
	panic(fmt.Errorf("not implemented: Requests - requests"))
}

// ScannedProducts is the resolver for the scannedProducts field.
func (r *queryResolver) ScannedProducts(ctx context.Context, requestID string) ([]*model.ScannedProductInfo, error) {
	panic(fmt.Errorf("not implemented: ScannedProducts - scannedProducts"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
