package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtSerivce interface {
	GenerateToken(userID uuid.UUID) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtCustomClaim struct {
	UserID uuid.UUID `json:"userID"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
}

func (js *jwtService) GenerateToken(userID uuid.UUID) string {
	claim := &JwtCustomClaim{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (js *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(js.secretKey), nil
	})
}

func NewJwtService(secretKey string) JwtSerivce {
	return &jwtService{
		secretKey: secretKey,
	}
}
