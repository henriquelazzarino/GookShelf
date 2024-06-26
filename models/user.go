package models

import (
	hashing "github.com/henriquelazzarino/gookshelf/utils"
)

type User struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	BookedBooks []Book `json:"bookedBooks"`
}

func (u *User) HashPassword() string {
	hash, err := hashing.HashPassword(u.Password)
	if err != nil {
		return ""
	}
	return hash
}

type UserRole string

const (
	Admin     UserRole = "admin"
	Librarian UserRole = "librarian"
	Regular   UserRole = "regular"
)
