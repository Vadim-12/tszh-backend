package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Port string
	DSN  string
}

func Load() Config {
	mode := os.Getenv("APP_ENV")
	envFilename := ".env." + mode

	if err := godotenv.Load(envFilename); err != nil {
		log.Fatalf("Failed to load env file: %s", envFilename)
	}

	log.Printf("Loaded env file: %s", envFilename)
	log.Println("DB_DSN =", os.Getenv("DB_DSN"))

	return Config{
		Env:  mode,
		Port: getenv("PORT", "8080"),
		DSN:  os.Getenv("DB_DSN"),
	}
}

func getenv(k, fallback string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return fallback
}
