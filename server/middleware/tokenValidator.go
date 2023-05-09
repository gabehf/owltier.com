package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/token"
)

type Values struct {
	m map[string]interface{}
}

func (v Values) GetUser() (*db.UserSchema, error) {
	u, ok := v.m["user"].(*db.UserSchema)
	if !ok || u == nil {
		return nil, errors.New("user is not set")
	}
	return u, nil
}
func (v Values) GetAccessToken() (string, error) {
	u, ok := v.m["access"].(string)
	if !ok {
		return "", errors.New("access token is not set")
	}
	return u, nil
}
func (v Values) GetRefreshToken() (string, error) {
	u, ok := v.m["refresh"].(string)
	if !ok {
		return "", errors.New("refresh token is not set")
	}
	return u, nil
}

func TokenValidater(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		headerVals := strings.Split(header, " ")
		if strings.ToLower(headerVals[0]) != "bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Bad Authorization Scheme")
		}
		t := headerVals[1]
		if t == "" {
			w.WriteHeader(401)
			fmt.Fprint(w, "No Token Provided")
			return
		}

		claims, err := token.ValidateAccess(t)
		if err != nil {
			w.WriteHeader(401)
			fmt.Fprint(w, "Unauthorized")
			return
		}

		if claims.Type != token.TypeAccess {
			w.WriteHeader(401)
			fmt.Fprint(w, "Unauthorized")
			return
		}

		var user = &db.UserSchema{}
		err = db.Fetch(&db.UserSchema{Id: claims.Id}, user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			fmt.Fprint(w, "User Not Found With Id: "+claims.Id)
			return
		}

		v := Values{map[string]interface{}{
			"user":   user,
			"access": t,
		}}

		ctx := context.WithValue(r.Context(), ContextKeyValues, &v)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func RefreshValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		headerVals := strings.Split(header, " ")
		if strings.ToLower(headerVals[0]) != "bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Bad Authorization Scheme")
		}
		t := headerVals[1]
		if t == "" {
			w.WriteHeader(401)
			fmt.Fprint(w, "No Token Provided")
			return
		}

		claims, err := token.ValidateRefresh(t)
		if err != nil {
			w.WriteHeader(401)
			fmt.Fprint(w, "Unauthorized")
			return
		}

		if claims.Type != token.TypeRefresh {
			w.WriteHeader(401)
			fmt.Fprint(w, "Unauthorized")
			return
		}

		var user = &db.UserSchema{}
		err = db.Fetch(&db.UserSchema{Id: claims.Id}, user)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, "User Not Found")
			return
		}

		v := Values{map[string]interface{}{
			"user":    user,
			"refresh": t,
		}}

		ctx := context.WithValue(r.Context(), ContextKeyValues, &v)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
