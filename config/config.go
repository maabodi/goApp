package config

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	JWT_KEY     string
}

func InitConfiguration() Config {
	return Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "root",
		DB_NAME:     "goapp",
		DB_PORT:     "3306",
		DB_HOST:     "localhost",
		JWT_KEY:     "rahasia",
	}
}
