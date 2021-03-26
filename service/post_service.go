package service

import (
	"context"

	"github.com/favecode/reflect-core/entity"
	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/pkg/db"
	"github.com/favecode/reflect-core/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)


func AddPost(ctx context.Context, input *model.AddPostInput) (*model.Post, error) {
	db := db.Connect()
	token, _ := util.GetDataFromHeaderWithKey(ctx, "token")
	userData, _ := util.ParseToken(token)

	isUser := &entity.User{}

	db.Model(isUser).Where("id = ?", userData.Id).Returning("*").Select()
	
	if (len(isUser.Id) <= 0) {
		return nil, gqlerror.Errorf("UserId not valid")
	}
	post := &entity.Post{
		UserId: isUser.Id,
		Content: input.Content,
		Link: util.RandomString(12),
	}

	db.Model(post).Returning("*").Insert()

	return &model.Post{
		ID: post.Id,
		User: &model.User{
			ID: isUser.Id,
			Username: isUser.Username,
			Email: isUser.Email,
			Fullname:isUser.Fullname,
			CreatedAt: &isUser.CreatedAt,
			UpdatedAt: &isUser.UpdatedAt,
		},
		Content: post.Content,
		Link: post.Link,
		CreatedAt: &post.CreatedAt,
		UpdatedAt: &post.UpdatedAt,	
	}, nil
}