package service

import (
	"context"
	"time"

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
		return nil, gqlerror.Errorf("UserId is not valid")
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
			DeletedAt: isUser.DeletedAt,
		},
		Content: post.Content,
		Link: post.Link,
		CreatedAt: &post.CreatedAt,
		UpdatedAt: &post.UpdatedAt,	
		DeletedAt: post.DeletedAt,	
	}, nil
}

func GetPost(ctx context.Context, input *model.GetPostInput) (*model.Post, error) {
	db := db.Connect()
	post := &entity.Post{}
	user := &entity.User{}
	db.Model(post).Where("id = ?", input.PostID).Returning("*").Select()
	db.Model(user).Where("id = ?", input.PostID).Returning("*").Select()

	if(len(post.Id) <= 0) {
		return nil, gqlerror.Errorf("PostId is not valid")	
	}

	return &model.Post{
		ID: post.Id,
		User: &model.User{
			ID: user.Id,
			Username: user.Username,
			Email: user.Email,
			Fullname:user.Fullname,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Content: post.Content,
		Link: post.Link,
		CreatedAt: &post.CreatedAt,
		UpdatedAt: &post.UpdatedAt,	
		DeletedAt: post.DeletedAt,	
	}, nil
}

func GetPostsByUsername(ctx context.Context, input *model.GetPostsByUsernameInput) ([]*model.Post, error) {
	db := db.Connect()
	posts := &[]entity.Post{}
	user := &entity.User{}
	db.Model(user).Where("username = ?", input.Username).Returning("*").Select()

	db.Model(posts).Where("user_id = ?", user.Id).Returning("*").Select()

	result := make([]*model.Post, 0)

	for _, post := range *posts {
		result = append(result, &model.Post{
			ID: post.Id,
			User: &model.User{
				ID: user.Id,
				Username: user.Username,
				Email: user.Email,
				Fullname:user.Fullname,
				CreatedAt: &user.CreatedAt,
				UpdatedAt: &user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
			Content: post.Content,
			Link: post.Link,
			CreatedAt: &post.CreatedAt,
			UpdatedAt: &post.UpdatedAt,		
			DeletedAt: post.DeletedAt,		
		})
	}	

	return result , nil
}


func DeletePost(ctx context.Context, input *model.DeletePostInput) (*model.Post, error) {
	db := db.Connect()
	token, _ := util.GetDataFromHeaderWithKey(ctx, "token")
	userData, _ := util.ParseToken(token)

	isUser := &entity.User{}

	db.Model(isUser).Where("id = ?", userData.Id).Returning("*").Select()

	if (len(isUser.Id) <= 0) {
		return nil, gqlerror.Errorf("UserId is not valid")
	}

	DeletedAt := time.Now()

	post := &entity.Post{
		Id: input.PostID,
		DeletedAt: &DeletedAt,
	}

	db.Model(post).Set("deleted_at = ?deleted_at").Where("id = ?id").Returning("*").Update()

	return &model.Post{
		ID: post.Id,
		User: &model.User{
			ID: isUser.Id,
			Username: isUser.Username,
			Email: isUser.Email,
			Fullname:isUser.Fullname,
			CreatedAt: &isUser.CreatedAt,
			UpdatedAt: &isUser.UpdatedAt,
			DeletedAt: isUser.DeletedAt,
		},
		Content: post.Content,
		Link: post.Link,
		CreatedAt: &post.CreatedAt,
		UpdatedAt: &post.UpdatedAt,	
		DeletedAt: post.DeletedAt,	
	} , nil
}