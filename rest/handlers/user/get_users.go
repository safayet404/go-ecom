package user

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.ListUser(), 200)

}
