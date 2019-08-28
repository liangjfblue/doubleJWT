package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

type Context struct {
	ID 			uint	`json:"id"`
	Username 	string	`json:"username"`
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

func ParseRequest(c *gin.Context, shortOrLong bool) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	var secret string
	if shortOrLong {
		secret = viper.GetString("jwt_secret.short")
	} else {
		secret = viper.GetString("jwt_secret.long")
	}


	if len(header) == 0 {
		return &Context{}, errors.New("`Authorization` header is 0")
	}

	var t string
	_, _ = fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

func SignToken(c Context) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":		c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":		time.Now().Add(time.Second * time.Duration(viper.GetInt("jwt_secret.shortTime"))).Unix(),	//use for request
	})

	tokenString, err = token.SignedString([]byte(viper.GetString("jwt_secret.short")))
	return
}

func SignRefreshToken(c Context) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":		c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":		time.Now().Add(time.Second * time.Duration(viper.GetInt("jwt_secret.longTime"))).Unix(),//use for refresh request token
	})

	tokenString, err = token.SignedString([]byte(viper.GetString("jwt_secret.long")))
	return
}