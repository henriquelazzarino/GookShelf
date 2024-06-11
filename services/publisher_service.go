package services

import (
    "errors"
    "github.com/henriquelazzarino/gookshelf/models"
	"github.com/henriquelazzarino/gookshelf/repositories"
)

func CreatePublisher(publisher *models.Publisher) error {
    // Verifica se o editor já existe
    _, err := repositories.GetPublisher(publisher.ID)
    if err == nil {
        return errors.New("publisher with same ID already exists")
    }

    return repositories.CreatePublisher(publisher)
}

func GetAllPublishers() ([]models.Publisher, error) {
    return repositories.GetAllPublishers()
}

func GetPublisher(id string) (*models.Publisher, error) {
    return repositories.GetPublisher(id)
}

func UpdatePublisher(id string, newPublisher *models.Publisher) error {
    // Verifica se o editor existe
    _, err := repositories.GetPublisher(id)
    if err != nil {
        return err
    }

    return repositories.UpdatePublisher(id, newPublisher)
}

func DeletePublisher(id string) error {
    // Verifica se o editor existe
    _, err := repositories.GetPublisher(id)
    if err != nil {
        return err
    }

    return repositories.DeletePublisher(id)
}
