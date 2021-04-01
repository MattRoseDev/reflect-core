package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/reflect-core/graph/generated"
	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/service"
)

func (r *mutationResolver) Register(ctx context.Context, input *model.RegisterInput) (*model.AuthOutput, error) {
	return service.Register(ctx, input)
}

func (r *queryResolver) Login(ctx context.Context, input *model.LoginInput) (*model.AuthOutput, error) {
	return service.Login(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
