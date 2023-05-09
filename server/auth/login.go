package auth

import (
	"net/http"
	"time"

	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/jsend"
	"github.com/mnrva-dev/owltier.com/server/token"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var form = &RequestForm{}
	if err := form.Parse(r); err != nil {
		jsend.Error(w, "Failed to parse form body")
		return
	}

	// get user from DB
	var user = &db.UserSchema{}
	err := db.FetchByGsi(&db.UserSchema{
		Email: form.Email,
	}, user)
	if err != nil {
		jsend.Fail(w, 401, map[string]interface{}{
			"message": "Email or password is invalid"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		jsend.Fail(w, 401, map[string]interface{}{
			"message": "Email or password is invalid"})
		return
	}

	// prepare login information for the client
	accessT := token.GenerateAccess(user)
	refreshT := token.GenerateRefresh(user)
	db.Update(user, "refresh_token", refreshT)
	db.Update(user, "last_login_at", time.Now().Unix())
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
	jsend.Success(w, map[string]interface{}{
		"id": user.Id,
	})
}
