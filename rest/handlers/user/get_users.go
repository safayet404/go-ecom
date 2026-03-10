package user

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	userList, err := h.userRepo.List()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	util.SendData(w, userList, 200)

}
