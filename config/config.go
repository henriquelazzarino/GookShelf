package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var (
    JWTSecret string
    Port      string
    FirebaseURL string
)

func LoadEnv() {
    err := godotenv.Load("config/.env")
    if err != nil {
        log.Fatalf("Error loading .env file")
        log.Fatalf("Error: %v", err)
    }

    JWTSecret = os.Getenv("JWT_SECRET")
    Port = os.Getenv("PORT")
    FirebaseURL = os.Getenv("FIREBASE_URL")

    if JWTSecret == "" || Port == "" || FirebaseURL == "" {
        log.Fatalf("Required environment variables are not set")
    }
}
