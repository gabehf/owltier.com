package auth

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/jsend"
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
	err := db.Fetch(&db.UserSchema{
		Username: form.Username,
	}, user)
	if err != nil {
		jsend.Fail(w, 401, map[string]interface{}{
			"message": "Username or password is invalid"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		jsend.Fail(w, 401, map[string]interface{}{
			"message": "Username or password is invalid"})
		return
	}

	// prepare login information for the client
	session := uuid.NewString()
	db.Update(user, "session", session)
	db.Update(user, "last_login_at", time.Now().Unix())
	http.SetCookie(w, &http.Cookie{
		Name:     SESSION_COOKIE,
		Value:    session,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   true,
	})
	jsend.Success(w, map[string]interface{}{
		"username": user.Username,
	})
}
