package config

import (
	"log"
	"github.com/joho/godotenv"
)


func Init(env string) {
	if env == "test" {
		err := godotenv.Load("./.env.test")
		if err != nil {
			log.Fatal("Error loading .env.test file")
		}
	} else if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}


