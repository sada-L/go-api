package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
