package crypto

import (
	"crypto/ecdsa"
	"errors"
	"log"

	"github.com/baturalpk/photo-bucket/config"
	"github.com/golang-jwt/jwt/v4"
)

var ecdsaPrivateKey *ecdsa.PrivateKey
var ecdsaPublicKey *ecdsa.PublicKey
var errInit error

func init() {
	LoadES256KeysIntoMemory()
}

func LoadES256KeysIntoMemory() {
	privateKey := config.Get().ES256.PrivateKey
	publicKey := config.Get().ES256.PublicKey

	ecdsaPrivateKey, errInit = jwt.ParseECPrivateKeyFromPEM([]byte(privateKey))
	if errInit != nil {
		panic(errInit.Error())
	}

	ecdsaPublicKey, errInit = jwt.ParseECPublicKeyFromPEM([]byte(publicKey))
	if errInit != nil {
		panic(errInit.Error())
	}
	log.Println("ES256 keys are loaded into the memory")
}

func SignNewJWTWithClaims(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAndExtractJWTClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return ecdsaPublicKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
