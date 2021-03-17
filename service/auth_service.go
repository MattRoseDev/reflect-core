package service

import (
	"context"

	"github.com/favecode/reflect-core/graph/model"
)

func Register(ctx context.Context, input *model.RegisterInput) (*model.AuthOutput, error) {
	return &model.AuthOutput{
		Token: "token",
		User: &model.User{
			Username: input.Username,
			Password: input.Password,
			Email: input.Email,
		},
	}, nil
}

func Login(ctx context.Context, input *model.LoginInput) (*model.AuthOutput, error) {
	return &model.AuthOutput{
		Token: "token",
		User: &model.User{
			Username: input.Username,
			Password: input.Password,
		},
	}, nil
}