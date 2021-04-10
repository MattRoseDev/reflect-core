package util

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/favecode/reflect-core/entity"
	"github.com/favecode/reflect-core/pkg/db"
	"github.com/gin-gonic/gin"
)

var jwtSecret []byte

type Claims struct {
	Id string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(110 * time.Hour)

	claims := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "reflect",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("header")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func GetDataFromHeaderWithKey(ctx context.Context, key string) (string, error) {
	gc, _ := GinContextFromContext(ctx)
	return gc.Request.Header.Get(key), nil
}

func ValidateUserToken(ctx context.Context) (*entity.User, error) {
	db := db.DB
	token, _ := GetDataFromHeaderWithKey(ctx, "token")
	userData, _ := ParseToken(token)

	if (userData == nil) {
		return nil, fmt.Errorf("token is not valid")
	}

	isUser := &entity.User{}

	db.Model(isUser).Where("id = ?", userData.Id).Where("deleted_at is ?", nil).Returning("*").Select()

	if (len(isUser.Id) <= 0) {
		return nil, fmt.Errorf("UserId is not valid")
	}

	return isUser, nil
}