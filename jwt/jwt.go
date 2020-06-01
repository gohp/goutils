package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	defaultKey      = "12345678"
	fifteenDayHours = 15 * 24
)

type Config struct {
	Key        string
	ExpireHour time.Duration
}

type CustomClaims struct {
	ID      interface{} // token unique id
	LoginAt int64       // login timestamp
	jwt.StandardClaims
}

type Service struct {
	C *Config
}

// New a token srv
func New(c *Config) *Service {
	if c == nil {
		c = &Config{
			Key:        defaultKey,
			ExpireHour: fifteenDayHours,
		}
	}
	s := &Service{
		C: c,
	}
	if c.ExpireHour == 0 {
		s.C.ExpireHour = fifteenDayHours
	}
	return s
}

// Decode a token string into a token object
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

// Encode a claim into a JWT
func (s *Service) EncodeToken(Id interface{}, loginAt time.Time) (string, error) {

	expireToken := time.Now().Add(time.Hour * s.C.ExpireHour).Unix()

	// Create the Claims
	claims := CustomClaims{
		Id,
		loginAt.Unix(),
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "mask.srv.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString([]byte(s.C.Key))
}
