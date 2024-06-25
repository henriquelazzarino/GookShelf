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
