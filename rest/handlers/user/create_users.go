package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {

	var newUsers ReqUser

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUsers)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give valid json", 400)
		return
	}

	user, err := h.userRepo.Create(repo.User{
		FirstName:   newUsers.FirstName,
		LastName:    newUsers.LastName,
		Password:    newUsers.Password,
		Email:       newUsers.Email,
		IsShopOwner: newUsers.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Internal Server error ", http.StatusInternalServerError)
		return
	}

	util.SendData(w, user, 201)

}
