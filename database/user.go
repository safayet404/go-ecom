package database

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

var userList []User

func UserStore(u User) User {

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

func DeleteUser(uId int) {
	var tempUser []User

	for _, user := range userList {
		if user.ID != uId {
			tempUser = append(tempUser, user)
		}
	}

	userList = tempUser
}

func init() {
	user1 := User{
		ID:     1,
		Name:   "Safayet Hossain",
		Email:  "hossainsafayet187@gmail.com",
		Mobile: "01679175553",
	}

	user2 := User{
		ID:     1,
		Name:   "Safayet Hossain",
		Email:  "hossainsafayet187@gmail.com",
		Mobile: "01679175553",
	}

	userList = append(userList, user1)
	userList = append(userList, user2)
}
