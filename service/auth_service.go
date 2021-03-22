package service

import (
	"context"

	"github.com/favecode/reflect-core/entity"
	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/pkg/db"
	"github.com/favecode/reflect-core/pkg/util"
	"github.com/google/uuid"
)

func Register(ctx context.Context, input *model.RegisterInput) (*model.AuthOutput, error) {
	db := db.Connect()
	id := uuid.NewString()
	user := &entity.User{
		Username: input.Username,
		Email: input.Email,
		Admin: false,
	}
	db.Model(user).Returning("*").Insert()
	password := &entity.Password{
		UserId: user.Id,
		Password: input.Password,
	}	
	db.Model(password).Insert()
	token, _ := util.GenerateToken(id, input.Username)
	return &model.AuthOutput{
		Token: token,
		User: &model.User{
			ID: user.Id,
			Username: user.Username,
			Email: user.Email,
			Fullname: user.Fullname,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		},
	}, nil
}

func Login(ctx context.Context, input *model.LoginInput) (*model.AuthOutput, error) {
	return &model.AuthOutput{
		Token: "token",
		User: &model.User{
			Username: input.Username,
		},
	}, nil
}