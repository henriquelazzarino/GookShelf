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

func AddBookToUser(user *models.User, book *models.Book) error {
	// Verifica se o livro já está na lista de livros do usuário
	for _, b := range user.BookedBooks {
		if b.ID == book.ID {
			return errors.New("book already added to user")
		}
	}

	if !book.IsFree {
		return errors.New("book is not available")
	}

	if book.MinimumAge > user.Age {
		return errors.New("user is not old enough to book this book")
	}

	// Adiciona o livro à lista de livros do usuário
	user.BookedBooks = append(user.BookedBooks, *book)

	// Atualiza o usuário no Firebase
	if err := UpdateUser(user.ID, user, false); err != nil {
		return err
	}

	SetIsFree(book.ID, false)

	return nil
}

func RemoveBookFromUser(user *models.User, book *models.Book) error {
	// Verifica se o livro está na lista de livros do usuário
	var bookIndex = -1
	for i, b := range user.BookedBooks {
		if b.ID == book.ID {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		return errors.New("book not found in user's list")
	}

	// Remove o livro da lista de livros do usuário
	user.BookedBooks = append(user.BookedBooks[:bookIndex], user.BookedBooks[bookIndex+1:]...)

	// Atualiza o usuário no Firebase
	if err := UpdateUser(user.ID, user, false); err != nil {
		return err
	}

	SetIsFree(book.ID, true)

	return nil
}

func UpdateUser(id string, newUser *models.User, hash bool) error {
	ref := config.FirebaseClient.NewRef("users").Child(id)
	var existingUser models.User
	if err := ref.Get(context.Background(), &existingUser); err != nil {
		return errors.New("user not found")
	}

	// Atualize a senha do usuário se necessário
	if hash {
		newUser.Password = newUser.HashPassword()
	}

	newUser.ID = existingUser.ID

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
