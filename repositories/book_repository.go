package repositories

import (
    "errors"
    "github.com/henriquelazzarino/gookshelf/models"
)

var books []models.Book

func CreateBook(book *models.Book) error {
    // Lógica para verificar se o livro já existe, por exemplo, verificando se o ID já está em uso
    for _, b := range books {
        if b.ID == book.ID {
            return errors.New("book with same ID already exists")
        }
    }

    // Adiciona o livro ao slice de livros
    books = append(books, *book)
    return nil
}

func GetAllBooks() ([]models.Book, error) {
    return books, nil
}

func GetBook(id string) (*models.Book, error) {
    // Procura pelo livro com o ID especificado
    for _, book := range books {
        if book.ID == id {
            return &book, nil
        }
    }
    return nil, errors.New("book not found")
}

func UpdateBook(id string, newBook *models.Book) error {
    // Procura pelo livro com o ID especificado
    for i, book := range books {
        if book.ID == id {
            // Atualiza os dados do livro
            books[i] = *newBook
            return nil
        }
    }
    return errors.New("book not found")
}

func DeleteBook(id string) error {
    // Procura pelo livro com o ID especificado
    for i, book := range books {
        if book.ID == id {
            // Remove o livro do slice de livros
            books = append(books[:i], books[i+1:]...)
            return nil
        }
    }
    return errors.New("book not found")
}
