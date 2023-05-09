package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mnrva-dev/owltier.com/server/config"
	"github.com/mnrva-dev/owltier.com/server/db"
)

type Claims struct {
	jwt.RegisteredClaims
	Id              string   `json:"id"`
	Username        string   `json:"username,omitempty"`
	Email           string   `json:"email"`
	EmailIsVerified bool     `json:"email_verified"`
	XSRF            string   `json:"xsrf"`
	Role            string   `json:"role,omitempty"`
	Policies        []string `json:"policies,omitempty"`
	Type            string   `json:"type"`
	Scope           string   `json:"scope"`
}

func BuildClaims(user *db.UserSchema) *Claims {
	var c = &Claims{
		Id:              user.Id,
		Username:        user.Username,
		Email:           user.Email,
		EmailIsVerified: user.EmailIsVerified,
		XSRF:            uuid.New().String(),
		Scope:           user.Scope,
		Policies:        user.Policies,
	}
	c.IssuedAt = jwt.NewNumericDate(time.Now())
	c.NotBefore = jwt.NewNumericDate(time.Now())
	c.Issuer = config.JwtIssuer()
	c.Audience = config.JwtAudience()
	if c.Scope == "" {
		c.Scope = "default"
	}
	return c
}
