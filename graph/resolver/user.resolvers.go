package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/service"
)

func (r *queryResolver) GetUserInfo(ctx context.Context) (*model.User, error) {
	return service.GetUserInfo(ctx)
}
