package repositories

import (
	"context"
	"errors"

	"github.com/henriquelazzarino/gookshelf/config"
	"github.com/henriquelazzarino/gookshelf/models"
)

func CreateUser(user *models.User) (string, error) {
	// Referência ao local no Firebase onde os usuários são armazenados
	ref := config.FirebaseClient.NewRef("users")

	// Adiciona o usuário ao Firebase com um ID gerado automaticamente
	newRef, err := ref.Push(context.Background(), user)
	if err != nil {
		return "", err
	}

	// Obtém o ID gerado pelo Firebase
	user.ID = newRef.Key

	// Atualiza o usuário com o ID gerado no Firebase
	if err := newRef.Update(context.Background(), map[string]interface{}{
		"id":       user.ID,
		"password": user.HashPassword(),
	}); err != nil {
		return "", err
	}

	return user.ID, nil
}

func GetAllUsers() ([]models.User, error) {
	ref := config.FirebaseClient.NewRef("users")
	var users map[string]models.User
	if err := ref.Get(context.Background(), &users); err != nil {
		return nil, err
	}

	var userList []models.User
	for _, user := range users {
		userList = append(userList, user)
	}

	return userList, nil
}

func GetUser(id string) (*models.User, error) {
	ref := config.FirebaseClient.NewRef("users").Child(id)
	var user models.User
	if err := ref.Get(context.Background(), &user); err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	ref := config.FirebaseClient.NewRef("users")
	var users map[string]models.User
	if err := ref.Get(context.Background(), &users); err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func UpdateUser(id string, newUser *models.User) error {
	ref := config.FirebaseClient.NewRef("users").Child(id)
	var existingUser models.User
	if err := ref.Get(context.Background(), &existingUser); err != nil {
		return errors.New("user not found")
	}

	// Atualize os dados do usuário
	if err := ref.Set(context.Background(), newUser); err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	ref := config.FirebaseClient.NewRef("users").Child(id)
	var existingUser models.User
	if err := ref.Get(context.Background(), &existingUser); err != nil {
		return errors.New("user not found")
	}

	// Remove o usuário do Firebase
	if err := ref.Delete(context.Background()); err != nil {
		return err
	}

	return nil
}
