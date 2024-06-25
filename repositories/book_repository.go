package repositories

import (
    "context"
    "errors"
    "github.com/henriquelazzarino/gookshelf/config"
    "github.com/henriquelazzarino/gookshelf/models"
)

func CreateBook(book *models.Book) (string, error) {
    // Referência ao local no Firebase onde os livros são armazenados
    ref := config.FirebaseClient.NewRef("books")

    // Adiciona o livro ao Firebase com um ID gerado automaticamente
    newRef, err := ref.Push(context.Background(), book)
    if err != nil {
        return "", err
    }

    // Obtém o ID gerado pelo Firebase
    book.ID = newRef.Key

    if err := newRef.Update(context.Background(), map[string]interface{}{
        "id": book.ID,
    }); err != nil {
        return "", err
    }

    return book.ID, nil
}

func GetAllBooks() ([]models.Book, error) {
    ref := config.FirebaseClient.NewRef("books")
    var books []models.Book
    if err := ref.Get(context.Background(), &books); err != nil {
        return nil, err
    }
    return books, nil
}

func GetBook(id string) (*models.Book, error) {
    ref := config.FirebaseClient.NewRef("books").Child(id)
    var book models.Book
    if err := ref.Get(context.Background(), &book); err != nil {
        return nil, errors.New("book not found")
    }
    return &book, nil
}

func UpdateBook(id string, newBook *models.Book) error {
    ref := config.FirebaseClient.NewRef("books").Child(id)
    if err := ref.Set(context.Background(), newBook); err != nil {
        return err
    }
    return nil
}

func DeleteBook(id string) error {
    ref := config.FirebaseClient.NewRef("books").Child(id)
    if err := ref.Delete(context.Background()); err != nil {
        return err
    }
    return nil
}
