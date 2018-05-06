package root

type User struct {
	Id string `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=50,username"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type UserService interface {
	CreateUser(u *User) error
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Login(c Credentials) (User, error)
}
