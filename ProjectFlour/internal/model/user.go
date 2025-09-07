package model

import "fmt"

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	if u.Name == "" || u.Username == "" || u.Password == "" {
		return fmt.Errorf("invalid input data")
	}

	return nil
}
