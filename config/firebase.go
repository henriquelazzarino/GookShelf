package config

import (
    "context"
    "firebase.google.com/go"
    "firebase.google.com/go/db"
    "log"

    "google.golang.org/api/option"
)

var FirebaseClient *db.Client

func InitFirebase() {
    // Usando as credenciais do serviço
    opt := option.WithCredentialsFile("config/serviceCredentials.json")

    // Inicializando a app Firebase
    app, err := firebase.NewApp(context.Background(), &firebase.Config{
        DatabaseURL: FirebaseURL,
    }, opt)
    if err != nil {
        log.Fatalf("erro ao inicializar o app Firebase: %v\n", err)
    }

    // Inicializando o cliente Realtime Database
    FirebaseClient, err = app.Database(context.Background())
    if err != nil {
        log.Fatalf("erro ao inicializar o cliente do Realtime Database: %v\n", err)
    }
}
