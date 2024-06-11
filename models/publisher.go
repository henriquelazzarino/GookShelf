package models

type Publisher struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Books []Book `json:"books"`
}
