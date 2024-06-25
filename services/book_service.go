package services

import (
    "errors"
    "github.com/henriquelazzarino/gookshelf/models"
	"github.com/henriquelazzarino/gookshelf/repositories"
)

func CreateBook(book *models.Book) (string, error) {
    // Verifica se o livro já existe
    _, err := repositories.GetBook(book.ID)
    if err != nil {
        return "", errors.New("book with same ID already exists")
    }

    return repositories.CreateBook(book)
}

func GetAllBooks() ([]models.Book, error) {
    return repositories.GetAllBooks()
}

func GetBook(id string) (*models.Book, error) {
    return repositories.GetBook(id)
}

func UpdateBook(id string, newBook *models.Book) error {
    // Verifica se o livro existe
    _, err := repositories.GetBook(id)
    if err != nil {
        return err
    }

    return repositories.UpdateBook(id, newBook)
}

func DeleteBook(id string) error {
    // Verifica se o livro existe
    _, err := repositories.GetBook(id)
    if err != nil {
        return err
    }

    return repositories.DeleteBook(id)
}
