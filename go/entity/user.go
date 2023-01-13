package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (u *User) ComparePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}
