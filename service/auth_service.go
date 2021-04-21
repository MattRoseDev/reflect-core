package service

import (
	"context"

	"github.com/favecode/reflect-core/entity"
	"github.com/favecode/reflect-core/graph/model"
	"github.com/favecode/reflect-core/pkg/db"
	"github.com/favecode/reflect-core/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx context.Context, input *model.RegisterInput) (*model.AuthOutput, error) {
	user := &entity.User{
		Username: util.RandomString(12),
		Email: input.Email,
		Fullname: &input.Fullname,
	}
	isUser := &entity.User{}

	db.DB.Model(isUser).Where("email = ?", input.Email).Where("deleted_at is ?", nil).Returning("*").Select()	

	if (isUser.Email == input.Email) {
		return nil, gqlerror.Errorf("Email has already taken")
	}
	
	db.DB.Model(user).Returning("*").Insert()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	hashedPassword := string(bytes)
	password := &entity.Password{
		UserId: user.Id,
		Password: hashedPassword,
	}	
	db.DB.Model(password).Insert()
	token, _ := util.GenerateToken(user.Id, user.Username)
	return &model.AuthOutput{
		Token: token,
		User: &model.User{
			ID: user.Id,
			Username: user.Username,
			Email: user.Email,
			Fullname: user.Fullname,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
	}, nil
}

func Login(ctx context.Context, input *model.LoginInput) (*model.AuthOutput, error) {
	user := new(entity.User)
	password := new(entity.Password)

	db.DB.Model(user).Where("username = ?", input.Username).WhereOr("email = ?", input.Username).Where("deleted_at is ?", nil).Select()
	if (len(user.Id) <= 0) {
		return nil, gqlerror.Errorf("Username or Password is not valid")
	}

	db.DB.Model(password).Where("user_id = ?", user.Id).Where("deleted_at is ?", nil).Select()	
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
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
		}, nil
	}
	return nil, gqlerror.Errorf("Username or Password is not valid")
}