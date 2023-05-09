package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/jsend"
	"github.com/mnrva-dev/owltier.com/server/middleware"
	"github.com/mnrva-dev/owltier.com/server/token"
)

func Refresh(w http.ResponseWriter, r *http.Request) {

	// get user and token from token parse middleware
	user, err := r.Context().Value(middleware.ContextKeyValues).(*middleware.Values).GetUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := r.Context().Value(middleware.ContextKeyValues).(*middleware.Values).GetRefreshToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Refresh != t {
		http.Error(w, "Token mismatch", http.StatusUnauthorized)
		fmt.Printf("%s\n%s", user.Refresh, t)
		return
	}

	// prepare login information for the client
	accessT := token.GenerateAccess(user)
	refreshT := token.GenerateRefresh(user)
	db.Update(user, "RefreshToken", refreshT)

	http.SetCookie(w, &http.Cookie{
		Name:     "_owltier.com_auth",
		Value:    accessT,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "_owltier.com_refresh",
		Value:    accessT,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   true,
	})
	jsend.Success(w, nil)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	headerVals := strings.Split(header, " ")
	if strings.ToLower(headerVals[0]) != "bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Bad Authorization Scheme")
	}
	t := headerVals[1]
	if t == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "No Token Provided")
		return
	}

	claims, err := token.ValidateAccess(t)
	if err != nil {
		w.WriteHeader(401)
		fmt.Fprint(w, "Unauthorized")
		return
	}

	if claims.Type != "Access" {
		w.WriteHeader(401)
		fmt.Fprint(w, "Unauthorized")
		return
	}
	jsend.Success(w, nil)
}
