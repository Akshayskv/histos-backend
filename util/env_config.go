package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type environmentVariables struct {
	DATABASE_URL          string
	JWT_SIGNING_ALGORITHM string
	JWT_SIGNING_SECRET    string
}

var EnvironmentVariables = loadEnvironmentVars()

func loadEnvironmentVars() environmentVariables {
	err := godotenv.Load()
	var env environmentVariables
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	env.DATABASE_URL = os.Getenv("DATABASE_URL")
	if env.DATABASE_URL == "" {
		log.Fatal("required environment variable missing")
	}

	env.JWT_SIGNING_ALGORITHM = os.Getenv("JWT_SIGNING_ALGORITHM")
	if env.JWT_SIGNING_ALGORITHM == "" {
		log.Fatal("required environment variable missing")
	}

	env.JWT_SIGNING_SECRET = os.Getenv("JWT_SIGNING_SECRET")
	if env.JWT_SIGNING_SECRET == "" {
		log.Fatal("required environment variable missing")
	}

	return env
}
