package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var newUser ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid User", 404)
		return
	}

	user := database.Find(newUser.Email, newUser.Password)

	if user == nil {
		http.Error(w, "Invalid credentials", 400)
		return
	}

	util.SendData(w, user, 200)

}
