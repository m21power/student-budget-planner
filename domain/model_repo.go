package domain

import "github.com/lib/pq"

type UserModel struct {
	ID       int            `json:"id" db:"id"`
	Name     string         `json:"name" db:"name"`
	Email    string         `json:"email" db:"email"`
	Password string         `json:"password" db:"password"`
	Badge    pq.StringArray `json:"badge" db:"badge"`
}

type UserRepository interface {
	Login(email, password string) (UserModel, error)
	Register(name, email, password string) error
	GetUser(id int) (UserModel, error)
	UpdateBadge(id int, badge string) error
}
