package config

type apiKeys struct {
	Telegram string
}

type Database struct {
	Name     string
	User     string
	Password string
}

type Config struct {
	ApiKeys apiKeys
	DB      Database
}
