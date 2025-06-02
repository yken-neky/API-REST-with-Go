package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type env struct{}

func loadEnv() *env {
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: .env file not found")
	}
	return &env{}
}

func (e *env) GetString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func (e *env) GetInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func (e *env) GetBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultValue
}
