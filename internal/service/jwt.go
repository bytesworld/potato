package service

import (
	"fmt"
	"gitee.com/bytesworld/tomato/configs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtService struct {
}

var JwtService = new(jwtService)

type JwtUser interface {
	GetUid() string
}
type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType    = "bearer"
	AppGuardName = "potato"
)

type TokenOutput struct {
	AccessToken string `json:"access_token,omitempty" yaml:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (js *jwtService) CreateToken(GuardName string, user JwtUser) (TokenOutput, error, *jwt.Token) {
	fmt.Println(user.GetUid())
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + configs.AppObj.Config.Jwt.JwtTtl+10000000,
				Id:        user.GetUid(),
				Issuer:    GuardName,
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(configs.AppObj.Config.Jwt.Secret))
	if err != nil {
		return TokenOutput{}, err, token
	}
	TokenData := TokenOutput{
		AccessToken: tokenStr,
		ExpiresIn:   configs.AppObj.Config.Jwt.JwtTtl+10,
		TokenType:   TokenType,
	}
	return TokenData, nil, token
}
