package list

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/jsend"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	// get iist from db
	listId := chi.URLParam(r, "id")
	list := &List{}
	list.Id = listId
	err := db.Fetch(list, list)
	if err != nil {
		jsend.Fail(w, 404, map[string]interface{}{
			"id": "list not found",
		})
		return
	}

	// return list
	jsend.Success(w, map[string]interface{}{
		"list": list,
	})
}
