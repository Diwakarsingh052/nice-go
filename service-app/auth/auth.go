package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

//create a key for the context

// Claims is our payload/core for out jwt token
type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

// Auth struct privateKey field would be used to verify and generate token
type Auth struct {
	privateKey *rsa.PrivateKey // this is key we would get after parsing the private.pem file
	publicKey  *rsa.PublicKey
}

func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Auth, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("private/public key cannot be nil")
	}
	a := Auth{
		privateKey: privateKey,
		publicKey:  publicKey,
	}

	return &a, nil

}

func (a *Auth) GenerateToken(claims Claims) (string, error) {

	//NewWithClaims creates a new Token with the specified signing method and claims
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	//signing our token with our private key
	tokenStr, err := tkn.SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}

	return tokenStr, nil

}

func (a *Auth) ValidateToken(tokenStr string) (Claims, error) {
	var c Claims
	token, err := jwt.ParseWithClaims(tokenStr, &c, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})

	if err != nil {
		return Claims{}, fmt.Errorf("parsing token %w", err)
	}
	if !token.Valid {
		return Claims{}, errors.New("invalid token")
	}

	//returning Claims if verification is successful
	return c, nil
}
