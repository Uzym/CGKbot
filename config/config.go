package config

import (
	"gopkg.in/yaml.v3"
)

type configFile struct {
	APIKeys struct {
		Telegram string `yaml:"telegram"`
	} `yaml:"apiKeys"`
	Database struct {
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {
	cf := configFile{}

	err := yaml.Unmarshal(fileBytes, &cf)
	if err != nil {
		return nil, err
	}

	c := Config{}

	c.ApiKeys.Telegram = cf.APIKeys.Telegram
	c.DB.Name = cf.Database.Name
	c.DB.User = cf.Database.User
	c.DB.Password = cf.Database.Password

	return &c, nil

}
