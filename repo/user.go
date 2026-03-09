package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(userID int) (*User, error)
	Find(email, pass string) (*User, error)
	List() ([]*User, error)
	Delete(userID int) error
	Update(p User) (*User, error)
}

type userRepo struct {
	userList []*User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}

	generateInitialUser(repo)

	return repo
}

func (r *userRepo) Create(u User) (*User, error) {
	u.ID = len(r.userList) + 1
	r.userList = append(r.userList, &u)

	return &u, nil
}
func (r *userRepo) Get(userID int) (*User, error) {
	for _, user := range r.userList {
		if user.ID == userID {

			return user, nil
		}
	}

	return nil, nil
}
func (r *userRepo) List() ([]*User, error) {
	return r.userList, nil
}
func (r *userRepo) Delete(userID int) error {
	var tempList []*User
	for _, u := range r.userList {
		if u.ID != userID {
			tempList = append(tempList, u)

		}
	}
	r.userList = tempList

	return nil
}
func (r *userRepo) Update(user User) (*User, error) {
	for idx, u := range r.userList {
		if u.ID == user.ID {
			r.userList[idx] = &user

		}
	}

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	for _, u := range r.userList {
		if u.Email == email && u.Password == pass {
			return u, nil
		}
	}

	return nil, nil

}

func generateInitialUser(r *userRepo) {
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

	r.userList = append(r.userList, &user1)
	r.userList = append(r.userList, &user2)

}
