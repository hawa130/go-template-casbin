package auth

import (
	"crypto/ecdsa"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hawa130/serverx/config"
	"github.com/rs/xid"
)

var (
	ErrInvalidToken = errors.New("jwt parse: invalid token")
)

type JWTClaims struct {
	Id        string `json:"jti"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
	Subject   xid.ID `json:"sub"`
}

// GenerateToken generates a new JWT token 为指定 ID 生成一个新的 JWT token
func GenerateToken(uid xid.ID) (string, error) {
	key, err := decodePem()
	if err != nil {
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"jti": xid.New().String(),
		"exp": time.Now().Add(config.Config().JWT.Duration * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"sub": uid.String(),
	})

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}

// ParseToken parses a JWT token 解析 JWT token
func ParseToken(token string) (*JWTClaims, error) {
	key, err := decodePem()
	if err != nil {
		return nil, err
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key.Public(), nil
	})
	if !t.Valid {
		return nil, ErrInvalidToken
	}
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	sub, err := xid.FromString(claims["sub"].(string))
	if err != nil {
		return nil, ErrInvalidToken
	}

	return &JWTClaims{
		Id:        claims["jti"].(string),
		IssuedAt:  int64(claims["iat"].(float64)),
		ExpiresAt: int64(claims["exp"].(float64)),
		Subject:   sub,
	}, nil

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
