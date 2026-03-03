package handlers

import (
	"ecommerce/config"
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

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShowOwner: user.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Internal Server error", 500)
	}

	util.SendData(w, accessToken, 200)

}
