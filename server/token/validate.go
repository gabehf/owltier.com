package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mnrva-dev/owltier.com/server/config"
	"golang.org/x/exp/slices"
)

func ValidateAccess(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.AccessSecret(), nil
	})

	if claims, ok := token.Claims.(*Claims); ok {
		if !token.Valid {
			return nil, fmt.Errorf("token is not valid")
		}
		if !slices.Contains(claims.Audience, "https://gosuimg.com") {
			return nil, fmt.Errorf("unexpected audience value: %v", claims.Audience)
		}
		if claims.Type != "Access" {
			return nil, fmt.Errorf("Unexpected token type: %v", claims.Type)
		}
		return claims, nil
	} else {
		return nil, err
	}
}

func ValidateRefresh(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.RefreshSecret(), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if !slices.Contains(claims.Audience, "https://gosuimg.com") {
			return &Claims{}, fmt.Errorf("unexpected audience value: %v", claims.Audience)
		}
		if claims.Type != "Refresh" {
			return &Claims{}, fmt.Errorf("Unexpected token type: %v", claims.Type)
		}
		return claims, nil
	} else {
		return &Claims{}, err
	}
}

func ValidateVerifyEmail(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.RefreshSecret(), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if !slices.Contains(claims.Audience, "https://gosuimg.com") {
			return &Claims{}, fmt.Errorf("unexpected audience value: %v", claims.Audience)
		}
		if claims.Type != "VerifyEmail" {
			return &Claims{}, fmt.Errorf("Unexpected token type: %v", claims.Type)
		}
		return claims, nil
	} else {
		return &Claims{}, err
	}
}
