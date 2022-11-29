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
		DB_USERNAME: "dikdik",
		DB_PASSWORD: "Kurniawan757!",
		DB_NAME:     "goapp",
		DB_PORT:     "3306",
		DB_HOST:     "178.128.209.164",
		JWT_KEY:     "rahasia",
	}
}
