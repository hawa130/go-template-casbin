package auth

import (
	"crypto/ecdsa"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hawa130/computility-cloud/config"
	"github.com/rs/xid"
)

var (
	ErrInvalidToken = errors.New("jwt parse: invalid token")
)

type JWTClaims struct {
	Uid xid.ID `json:"id"`
	jwt.StandardClaims
}

// GenerateToken generates a new JWT token
func GenerateToken(uid xid.ID) (string, error) {
	key, err := decodePem()
	if err != nil {
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodES256, JWTClaims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			Id:        xid.New().String(),
			IssuedAt:  jwt.TimeFunc().Unix(),
			ExpiresAt: jwt.TimeFunc().Add(config.Config().JWT.Duration * time.Hour).Unix(),
		},
	})

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}

// ParseToken parses a JWT token
func ParseToken(token string) (*JWTClaims, error) {
	key, err := decodePem()
	if err != nil {
		return nil, err
	}

	t, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key.Public(), nil
	})
	if !t.Valid {
		return nil, ErrInvalidToken
	}
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*JWTClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil

}

func decodePem() (*ecdsa.PrivateKey, error) {
	filePath := config.Config().JWT.PrivateKeyPath
	keyPem, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseECPrivateKeyFromPEM(keyPem)
	if err != nil {
		return nil, err
	}
	return key, nil
}
