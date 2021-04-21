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
	isUser, err := util.ValidateUserToken(ctx)	
	if (err != nil) {
		return nil, gqlerror.Errorf(err.Error())
	}

	post := &entity.Post{
		UserId: isUser.Id,
		Content: input.Content,
		Link: util.RandomString(12),
	}

	db.DB.Model(post).Returning("*").Insert()

	return &model.Post{
		ID: post.Id,
		User: &model.User{
			ID: isUser.Id,
			Username: isUser.Username,
			Email: isUser.Email,
			Fullname:isUser.Fullname,
			CreatedAt: isUser.CreatedAt,
			UpdatedAt: isUser.UpdatedAt,
			DeletedAt: isUser.DeletedAt,
		},
		Content: post.Content,
		Link: post.Link,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,	
		DeletedAt: post.DeletedAt,	
	}, nil
}

func GetPost(ctx context.Context, input *model.GetPostInput) (*model.Post, error) {
	_, err := util.ValidateUserToken(ctx)	
	if (err != nil) {
		return nil, gqlerror.Errorf(err.Error())
	}

	post := &entity.Post{}
	user := &entity.User{}
	db.DB.Model(post).Where("id = ?", input.PostID).Where("deleted_at is ?", nil).Returning("*").Select()
	db.DB.Model(user).Where("id = ?", input.PostID).Where("deleted_at is ?", nil).Returning("*").Select()

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
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Content: post.Content,
		Link: post.Link,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,	
		DeletedAt: post.DeletedAt,	
	}, nil
}

func GetPostsByUsername(ctx context.Context, input *model.GetPostsByUsernameInput) ([]*model.Post, error) {
	_, err := util.ValidateUserToken(ctx)	
	if (err != nil) {
		return nil, gqlerror.Errorf(err.Error())
	}

	posts := &[]entity.Post{}
	user := &entity.User{}
	db.DB.Model(user).Where("username = ?", input.Username).Where("deleted_at is ?", nil).Returning("*").Select()

	db.DB.Model(posts).Where("user_id = ?", user.Id).Order("created_at DESC").Returning("*").Select()

	result := make([]*model.Post, 0)

	for _, post := range *posts {
		result = append(result, &model.Post{
			ID: post.Id,
			User: &model.User{
				ID: user.Id,
				Username: user.Username,
				Email: user.Email,
				Fullname:user.Fullname,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
			Content: post.Content,
			Link: post.Link,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,		
			DeletedAt: post.DeletedAt,		
		})
	}	

	return result , nil
}

func DeletePost(ctx context.Context, input *model.DeletePostInput) (*model.Post, error) {
	isUser, err := util.ValidateUserToken(ctx)	
	if (err != nil) {
		return nil, gqlerror.Errorf(err.Error())
	}

	isPost := new(entity.Post)

	db.DB.Model(isPost).Where("id = ?", input.PostID).Returning("*").Select()
	if (len(isPost.Id) <= 0) {
		return nil, gqlerror.Errorf("Post not found")
	}
	
	DeletedAt := time.Now()

	post := &entity.Post{
		Id: input.PostID,
		DeletedAt: &DeletedAt,
	}
	db.DB.Model(post).Set("deleted_at = ?deleted_at").Where("id = ?id").Returning("*").Update()

	return &model.Post{
		ID: post.Id,
		User: &model.User{
			ID: isUser.Id,
			Username: isUser.Username,
			Email: isUser.Email,
			Fullname:isUser.Fullname,
			CreatedAt: isUser.CreatedAt,
			UpdatedAt: isUser.UpdatedAt,
			DeletedAt: isUser.DeletedAt,
		},
		Content: post.Content,
		Link: post.Link,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,	
		DeletedAt: post.DeletedAt,	
	} , nil
}

func EditPost(ctx context.Context, input *model.EditPostInput) (*model.Post, error) {
	isUser, err := util.ValidateUserToken(ctx)	
	if (err != nil) {
		return nil, gqlerror.Errorf(err.Error())
	}

	post := &entity.Post{
		Id: input.PostID,
		Content: input.Content,
	}

	db.DB.Model(post).Set("content = ?content").Where("id = ?id").Returning("*").Update()

	return &model.Post{
		ID: post.Id,
		User: &model.User{
			ID: isUser.Id,
			Username: isUser.Username,
			Email: isUser.Email,
			Fullname:isUser.Fullname,
			CreatedAt: isUser.CreatedAt,
			UpdatedAt: isUser.UpdatedAt,
			DeletedAt: isUser.DeletedAt,
		},
		Content: post.Content,
		Link: post.Link,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,	
		DeletedAt: post.DeletedAt,	
	} , nil
}