package list

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/mnrva-dev/owltier.com/server/auth"
	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/jsend"
)

const (
	ID_LENGTH = 10
)

func NewList(w http.ResponseWriter, r *http.Request) {

	// get user from session key
	sessC, err := r.Cookie(auth.SESSION_COOKIE)
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

	// create list object from user and json data
	list := &List{}
	json.NewDecoder(r.Body).Decode(list)
	list.CreatedAt = time.Now().Unix()
	list.CreatedBy = user.Username
	list.Id = makeId(ID_LENGTH)

	// create list in db
	db.Create(list)

	jsend.Success(w, map[string]interface{}{
		"slug": "/list/" + list.Id,
	})
}

func makeId(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
