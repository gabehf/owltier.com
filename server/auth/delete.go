package auth

import (
	"net/http"
	"strings"

	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/jsend"
	"github.com/mnrva-dev/owltier.com/server/middleware"
	"golang.org/x/crypto/bcrypt"
)

// Need login details AND valid token to delete an account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	password := strings.TrimSpace(r.FormValue("password"))

	// get user from token parse middleware
	user, err := r.Context().Value(middleware.ContextKeyValues).(*middleware.Values).GetUser()
	if err != nil {
		jsend.Error(w, "Failed to retrieve user information")
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
