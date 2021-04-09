package service

import (
	"context"

	"github.com/favecode/reflect-core/graph/model"
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