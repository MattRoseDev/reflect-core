package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/favecode/reflect-core/graph/generated"
	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/service"
)

func (r *mutationResolver) AddPost(ctx context.Context, input *model.AddPostInput) (*model.Post, error) {
	return service.AddPost(ctx, input)
}

func (r *mutationResolver) EditPost(ctx context.Context, input *model.EditPostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *model.DeletePostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPost(ctx context.Context, input *model.GetPostInput) (*model.Post, error) {
	return service.GetPost(ctx, input)
}

func (r *queryResolver) GetPostsByUsername(ctx context.Context, input *model.GetPostsByUsernameInput) ([]*model.Post, error) {
	return service.GetPostsByUsername(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
