package service

import (
	"context"

	"github.com/favecode/reflect-core/entity"
	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/pkg/db"
	"github.com/favecode/reflect-core/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GetUserInfo(ctx context.Context) (*model.User, error) {
	db := db.Connect()
	token, _ := util.GetDataFromHeaderWithKey(ctx, "token")
	userData, _ := util.ParseToken(token)

	isUser := &entity.User{}

	db.Model(isUser).Where("id = ?", userData.Id).Where("deleted_at is ?", nil).Returning("*").Select()

	if (len(isUser.Id) <= 0) {
		return nil, gqlerror.Errorf("UserId is not valid")
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