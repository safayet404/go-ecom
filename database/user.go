package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var userList []User

func UserStore(u User) User {

	if u.ID != 0 {
		return u
	}

	u.ID = len(userList) + 1
	userList = append(userList, u)
	return u

}

func ListUser() []User {
	return userList
}

func GetUser(uId int) *User {

	for _, user := range userList {
		if user.ID == uId {
			return &user
		}
	}
	return nil
}

func UpdateUser(user User) {
	for idx, u := range userList {
		if u.ID == user.ID {
			userList[idx] = user
		}
	}
}

func DeleteUser(uId, pass int) {
	var tempUser []User

	for _, user := range userList {
		if user.ID != uId {
			tempUser = append(tempUser, user)
		}
	}

	userList = tempUser
}

func Find(email, pass string) *User {
	for _, u := range userList {
		if u.Email == email && u.Password == pass {
			return &u
		}
	}

	return nil
}

func init() {
	user1 := User{
		ID:          1,
		FirstName:   "Safayet Hossain",
		LastName:    "Arif",
		Password:    "password123@sge.",
		Email:       "stgaraf72@gmail.com",
		IsShopOwner: true,
	}

	user2 := User{
		ID:          1,
		FirstName:   "Safayet Hossain",
		LastName:    "Araf",
		Email:       "hossainsafayet187@gmail.com",
		Password:    "password123@sge.",
		IsShopOwner: false,
	}

	userList = append(userList, user1)
	userList = append(userList, user2)
}
