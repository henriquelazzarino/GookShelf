package services

import (
	"errors"

	"github.com/henriquelazzarino/gookshelf/models"
	"github.com/henriquelazzarino/gookshelf/repositories"
)

func CreateUser(user *models.User) (string, error) {
	// Verifica se o usuário já existe
	_, err := repositories.GetUser(user.ID)
	if err != nil {
		return "", errors.New("user with same ID already exists")
	}

	return repositories.CreateUser(user)
}

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func GetUser(id string) (*models.User, error) {
	return repositories.GetUser(id)
}

func GetUserByEmail(email string) (*models.User, error) {
	return repositories.GetUserByEmail(email)
}

func AddBookToUser(userId string, bookId string) error {
	// Verifica se o usuário existe
	user, err := repositories.GetUser(userId)
	if err != nil {
		return err
	}

	// Verifica se o livro existe
	book, err := repositories.GetBook(bookId)
	if err != nil {
		return err
	}

	return repositories.AddBookToUser(user, book)
}

func RemoveBookFromUser(userId string, bookId string) error {
	// Verifica se o usuário existe
	user, err := repositories.GetUser(userId)
	if err != nil {
		return err
	}

	// Verifica se o livro existe
	book, err := repositories.GetBook(bookId)
	if err != nil {
		return err
	}

	return repositories.RemoveBookFromUser(user, book)
}

func UpdateUser(id string, newUser *models.User) error {
	// Verifica se o usuário existe
	_, err := repositories.GetUser(id)
	if err != nil {
		return err
	}

	return repositories.UpdateUser(id, newUser)
}

func DeleteUser(id string) error {
	// Verifica se o usuário existe
	_, err := repositories.GetUser(id)
	if err != nil {
		return err
	}

	return repositories.DeleteUser(id)
}
