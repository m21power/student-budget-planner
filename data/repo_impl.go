package data

import (
	"database/sql"
	"fmt"
	"student-planner/domain"

	"github.com/lib/pq"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetUser(id int) (domain.UserModel, error) {
	var user domain.UserModel
	err := s.db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Badge)
	if err != nil {
		return domain.UserModel{}, err
	}
	if user.Badge == nil {
		user.Badge = pq.StringArray{}
	}
	return user, nil
}

func (s *UserStore) Login(email string, password string) (domain.UserModel, error) {
	query := "SELECT * FROM users WHERE email=$1"
	user := s.db.QueryRow(query, email)
	var u domain.UserModel
	err := user.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Badge)
	if u.Badge == nil {
		u.Badge = pq.StringArray{} //set to empty array
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.UserModel{}, fmt.Errorf("user not found")
		}
		return domain.UserModel{}, err
	}
	if u.Password == password {
		return u, nil
	}
	return domain.UserModel{}, fmt.Errorf("password incorrect")
}

func (s *UserStore) Register(name string, email string, password string) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, name, email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) UpdateBadge(id int, badge string) error {
	query := "UPDATE users SET badges = array_append(badges, $1) WHERE id=$2"
	_, err := s.db.Exec(query, badge, id)
	if err != nil {
		return err
	}
	return nil
}
