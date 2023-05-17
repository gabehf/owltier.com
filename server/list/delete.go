package list

import (
	"net/http"

	"github.com/mnrva-dev/owltier.com/server/auth"
	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/jsend"
)

func DeleteList(w http.ResponseWriter, r *http.Request) {
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
	// get list id
	r.ParseForm()
	listId := r.FormValue("id")
	list := &List{Id: listId}
	// make sure user owns list
	err = db.Fetch(list, list)
	if err != nil {
		jsend.Fail(w, 404, map[string]interface{}{
			"id": "list not found",
		})
		return
	}
	if list.CreatedBy != user.Username {
		jsend.Fail(w, 401, map[string]interface{}{
			"username": "user does not own this list",
		})
		return
	}
	// delete list
	db.Delete(list)

	jsend.Success(w, map[string]interface{}{
		"id": list.Id,
	})
}
