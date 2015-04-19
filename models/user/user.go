package user

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
}

func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func (u *User) Save() {

}
