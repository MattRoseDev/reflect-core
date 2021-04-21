package service

import (
	"context"
	"fmt"

	"github.com/favecode/reflect-core/entity"
	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/pkg/db"
	"github.com/favecode/reflect-core/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GetUserInfo(ctx context.Context) (*model.User, error) {
	isUser, err := util.ValidateUserToken(ctx)	
	if (err != nil) {
		return nil, gqlerror.Errorf(err.Error())
	}

	return &model.User{
		ID: isUser.Id,
		Username: isUser.Username,
		Email: isUser.Email,
		Fullname: isUser.Fullname,
		Admin: isUser.Admin,
		Bio: isUser.Bio,
		CreatedAt: isUser.CreatedAt,
		UpdatedAt: isUser.UpdatedAt,
		DeletedAt: isUser.DeletedAt,
	} ,nil
}

func GetUserByUsername(ctx context.Context, input *model.GetUserByUsernameInput) (*model.User, error) {
	user := &entity.User{}
	db.DB.Model(user).Where("username = ?", input.Username).Where("deleted_at is ?", nil).Returning("*").Select() 

	if (len(user.Id) <= 0) {
		return nil, fmt.Errorf("user not found")
	}

	return &model.User{
		ID: user.Id,
		Username: user.Username,
		Email: user.Email,
		Fullname: user.Fullname,
		Admin: user.Admin,
		Bio: user.Bio,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	} ,nil
}