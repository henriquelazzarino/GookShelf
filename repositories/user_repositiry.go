package repositories

import (
    "errors"
    "github.com/henriquelazzarino/gookshelf/models"
)

var users []models.User

func CreateUser(user *models.User) error {
    // Lógica para verificar se o usuário já existe, por exemplo, verificando se o ID já está em uso
    for _, u := range users {
        if u.ID == user.ID {
            return errors.New("user with same ID already exists")
        }
    }

    // Adiciona o usuário ao slice de usuários
    users = append(users, *user)
    return nil
}

func GetAllUsers() ([]models.User, error) {
    return users, nil
}

func GetUser(id string) (*models.User, error) {
    // Procura pelo usuário com o ID especificado
    for _, user := range users {
        if user.ID == id {
            return &user, nil
        }
    }
    return nil, errors.New("user not found")
}

func UpdateUser(id string, newUser *models.User) error {
    // Procura pelo usuário com o ID especificado
    for i, user := range users {
        if user.ID == id {
            // Atualiza os dados do usuário
            users[i] = *newUser
            return nil
        }
    }
    return errors.New("user not found")
}

func DeleteUser(id string) error {
    // Procura pelo usuário com o ID especificado
    for i, user := range users {
        if user.ID == id {
            // Remove o usuário do slice de usuários
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return errors.New("user not found")
}
