package main

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID  string
	IsAdmin bool
}

func init() {
	publicKeyByte, err := ioutil.ReadFile("./rsa_public.pem")
	if err != nil {
		panic(err.Error())
	}
	publicKey, _ = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)

	privateKeyByte, err := ioutil.ReadFile("./rsa_private.pem")
	if err != nil {
		panic(err.Error())
	}
	privateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
}

func genToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func MakeToken() (string, error) {
	claims := &JWTClaims{
		UserID:  "120201080412",
		IsAdmin: false,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(3600)).Unix()
	return genToken(claims)
}

func VerifyToken(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("failed to change claim type")
	}
	if err = token.Claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}

func main() {
	strToken, _ := MakeToken()
	println(strToken)
	claim, _ := VerifyToken(strToken)
	fmt.Printf("%v", claim)
}
