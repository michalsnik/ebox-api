package config

import "github.com/ianschenck/envflag"

var (
	env         = envflag.String("ENV", "development", "Environment where the code is running. one of development, staging or production")

	dbHost 		= envflag.String("DB_HOST", "localhost", "Database Host")
	dbPort 		= envflag.Int("DB_PORT", 5434, "Database Port")
	dbUser 		= envflag.String("DB_USER", "postgres", "Database User")
	dbPassword 	= envflag.String("DB_PASSWORD", "supersecret", "Database Password")
	dbName 		= envflag.String("DB_NAME", "postgres", "Database Name")
)

type Config struct {
	DB *DBConfig
	App *AppConfig
}

type DBConfig struct {
	Host string
	Port int
	Username string
	Password string
	DBName string
}

type AppConfig struct {
	Env string
}

func GetConfig () *Config {
	return &Config{
		App: &AppConfig{
			Env: *env,
		},
		DB: &DBConfig{
			Host: *dbHost,
			Port: *dbPort,
			Username: *dbUser,
			Password: *dbPassword,
			DBName: *dbName,
		},
	}
}