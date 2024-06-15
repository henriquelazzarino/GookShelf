package models

type User struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    Email        string `json:"email"`
    Password     string `json:"password"`
    Picture      string `json:"picture"`
    Age          int    `json:"age"`
    BookedBooks  []Book `json:"bookedBooks"`
}

type UserRole string

const (
    Admin      UserRole = "admin"
    Librarian  UserRole = "librarian"
    Regular    UserRole = "regular"
)
