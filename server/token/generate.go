package token

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mnrva-dev/owltier.com/server/config"
	"github.com/mnrva-dev/owltier.com/server/db"
)

const (
	ACCESS_EXPIRATION       = 10 * time.Minute
	REFRESH_EXPIRATION      = 7 * 24 * time.Hour
	VERIFY_EMAIL_EXPIRATION = 15 * time.Minute
)

func GenerateAccess(user *db.UserSchema) string {
	c := BuildClaims(user)
	c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(ACCESS_EXPIRATION))
	c.Type = TypeAccess
	c.Id = user.Id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(config.AccessSecret())
	if err != nil {
		log.Println(err)
		return ""
	}

	return tokenString
}

func GenerateRefresh(user *db.UserSchema) string {
	c := BuildClaims(user)
	c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(REFRESH_EXPIRATION))
	c.Type = TypeRefresh
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(config.RefreshSecret())
	if err != nil {
		log.Println(err)
		return ""
	}

	return tokenString
}

func GenerateVerifyEmail(user *db.UserSchema) string {
	c := &Claims{}
	c.Id = user.Id
	c.Email = user.Email
	c.IssuedAt = jwt.NewNumericDate(time.Now())
	c.NotBefore = jwt.NewNumericDate(time.Now())
	c.Issuer = config.JwtIssuer()
	c.Audience = config.JwtAudience()
	c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(VERIFY_EMAIL_EXPIRATION))
	c.Type = TypeVerifyEmail
	c.Scope = "verify-email"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(config.EmailTokenSecret())
	if err != nil {
		log.Println(err)
		return ""
	}

	return tokenString
}
