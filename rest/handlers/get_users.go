package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.ListUser(), 200)

}
