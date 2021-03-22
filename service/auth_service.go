package service

import (
	"context"

	"github.com/favecode/reflect-core/entity"
	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/pkg/db"
	"github.com/favecode/reflect-core/pkg/util"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	bytes, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	hashedPassword := string(bytes)
	password := &entity.Password{
		UserId: user.Id,
		Password: hashedPassword,
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
	db := db.Connect()
	user := new(entity.User)
	password := new(entity.Password)
	db.Model(user).Where("username = ?", input.Username).Select()	
	db.Model(password).Where("user_id = ?", user.Id).Select()	
	err := bcrypt.CompareHashAndPassword([]byte(password.Password), []byte(input.Password)) 
	if (err == nil) {
		token, _ := util.GenerateToken(user.Id, user.Username)
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
	return nil, err
}