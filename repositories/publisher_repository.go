package repositories

import (
    "errors"
    "github.com/henriquelazzarino/gookshelf/models"
)

var publishers []models.Publisher

func CreatePublisher(publisher *models.Publisher) error {
    // Lógica para verificar se o editor já existe, por exemplo, verificando se o ID já está em uso
    for _, p := range publishers {
        if p.ID == publisher.ID {
            return errors.New("publisher with same ID already exists")
        }
    }

    // Adiciona o editor ao slice de editores
    publishers = append(publishers, *publisher)
    return nil
}

func GetAllPublishers() ([]models.Publisher, error) {
    return publishers, nil
}

func GetPublisher(id string) (*models.Publisher, error) {
    // Procura pelo editor com o ID especificado
    for _, publisher := range publishers {
        if publisher.ID == id {
            return &publisher, nil
        }
    }
    return nil, errors.New("publisher not found")
}

func UpdatePublisher(id string, newPublisher *models.Publisher) error {
    // Procura pelo editor com o ID especificado
    for i, publisher := range publishers {
        if publisher.ID == id {
            // Atualiza os dados do editor
            publishers[i] = *newPublisher
            return nil
        }
    }
    return errors.New("publisher not found")
}

func DeletePublisher(id string) error {
    // Procura pelo editor com o ID especificado
    for i, publisher := range publishers {
        if publisher.ID == id {
            // Remove o editor do slice de editores
            publishers = append(publishers[:i], publishers[i+1:]...)
            return nil
        }
    }
    return errors.New("publisher not found")
}
