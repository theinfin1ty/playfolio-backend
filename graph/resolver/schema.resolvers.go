package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"fmt"
	graph1 "playfolio-backend/graph"
)

// HealthMutation is the resolver for the healthMutation field.
func (r *mutationResolver) HealthMutation(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: HealthMutation - healthMutation"))
}

// HealthQuery is the resolver for the healthQuery field.
func (r *queryResolver) HealthQuery(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: HealthQuery - healthQuery"))
}

// Mutation returns graph1.MutationResolver implementation.
func (r *Resolver) Mutation() graph1.MutationResolver { return &mutationResolver{r} }

// Query returns graph1.QueryResolver implementation.
func (r *Resolver) Query() graph1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
