package config

import (
	"flag"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBname   string `yaml:"dbname"`
	} `yaml:"database"`
	Api struct {
		Port int `yaml:"port"`
	} `yaml:"api"`
}

func LoadConfig() (*Config, error) {
	env := flag.String("env", "dev", "environment immunity is running in")
	switch *env {
	case "dev":
		viper.SetConfigName("dev")
		viper.AddConfigPath("./config")
	case "production":
		viper.SetConfigName("immunity")
		viper.AddConfigPath("/etc/immunity/")
	}
	viper.SetConfigType("yaml")
	//viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
