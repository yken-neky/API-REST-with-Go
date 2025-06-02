package config

import (
	"fmt"
	"sync"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Server   ServerConfig
}

type AppConfig struct {
	Name        string
	Version     string
	Environment string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

var (
	config *Config
	once   sync.Once
)

// GetConfig retorna la configuración global de la aplicación
func GetConfig() *Config {
	once.Do(func() {
		env := loadEnv()
		config = &Config{
			App: AppConfig{
				Name:        env.GetString("APP_NAME", "api-rest-with-go"),
				Version:     env.GetString("APP_VERSION", "1.0.0"),
				Environment: env.GetString("APP_ENV", "development"),
			},
			Database: DatabaseConfig{
				Host:     env.GetString("DB_HOST", "localhost"),
				Port:     env.GetString("DB_PORT", "5432"),
				User:     env.GetString("DB_USER", "postgres"),
				Password: env.GetString("DB_PASSWORD", ""),
				DBName:   env.GetString("DB_NAME", "api_rest_go"),
				SSLMode:  env.GetString("DB_SSLMODE", "disable"),
			},
			Server: ServerConfig{
				Port:         env.GetString("SERVER_PORT", "8080"),
				ReadTimeout:  env.GetInt("SERVER_READ_TIMEOUT", 10),
				WriteTimeout: env.GetInt("SERVER_WRITE_TIMEOUT", 10),
				IdleTimeout:  env.GetInt("SERVER_IDLE_TIMEOUT", 60),
			},
		}
	})
	return config
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}
