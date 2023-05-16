package auth

import (
	"net/http"
	"strings"

	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/jsend"
	"golang.org/x/crypto/bcrypt"
)

// Need login details AND valid token to delete an account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	password := strings.TrimSpace(r.FormValue("password"))

	// get user session id from cookies
	sessC, err := r.Cookie(SESSION_COOKIE)
	if err != nil {
		jsend.Fail(w, 401, map[string]interface{}{
			"session": "invalid session",
		})
		return
	}
	session := sessC.Value
	var user = &db.UserSchema{}
	err = db.FetchByGsi(&db.UserSchema{
		Session: session,
	}, user)
	if err != nil {
		jsend.Fail(w, 401, map[string]interface{}{
			"session": "invalid session",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		jsend.Fail(w, 401, map[string]interface{}{
			"password": "Password is incorrect"})
		return
	}

	// at this point, password is correct and token is valid

	err = db.Delete(user)
	if err != nil {
		jsend.Error(w, "Failed to delete user")
		return
	}
	jsend.Success(w, nil)
}
