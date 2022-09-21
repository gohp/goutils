package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	defaultKey      = "12345678"
	fifteenDayHours = 15 * 24
)

type Config struct {
	Key        string
	ExpireTime time.Duration
	Issuer     string
}

type CustomClaims struct {
	ID      interface{} // token unique id
	LoginAt int64       // login timestamp
	jwt.RegisteredClaims
}

type Service struct {
	C *Config
}

// New a token srv
func New(c *Config) *Service {
	if c == nil {
		c = &Config{
			Key:        defaultKey,
			ExpireTime: time.Hour * fifteenDayHours,
		}
	}
	s := &Service{
		C: c,
	}
	if c.ExpireTime == 0 {
		s.C.ExpireTime = time.Hour * fifteenDayHours
	}
	return s
}

// DecodeToken a token string into a token object
func (s *Service) DecodeToken(tokenString string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.C.Key), nil
	})
	if err != nil {
		return nil, errors.New("")
	}
	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// EncodeToken a claim into a JWT
func (s *Service) EncodeToken(Id interface{}, loginAt time.Time) (string, error) {
	// Create the Claims
	claims := CustomClaims{
		Id,
		loginAt.Unix(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.C.ExpireTime)),
			Issuer:    s.C.Issuer,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString([]byte(s.C.Key))
}
