package user

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {

	var newUsers database.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUsers)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give valid json", 400)
		return
	}

	database.UserStore(newUsers)

	util.SendData(w, "New User Created", 201)

}
