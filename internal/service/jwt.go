package service

import "github.com/dgrijalva/jwt-go"

type jwtService struct {

}
var JwtService=new(jwtService)

type JwtUser interface {
	GetUid() string
}
type CustomClaims struct {
	jwt.StandardClaims
}