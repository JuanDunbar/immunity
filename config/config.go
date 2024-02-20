package config

import (
	"flag"
	"path/filepath"
	"runtime"

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
		workingDir := getWorkingDir()
		viper.AddConfigPath(workingDir)
		viper.SetConfigName("dev")
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

func getWorkingDir() string {
	_, b, _, _ := runtime.Caller(0)
	wDir := filepath.Dir(b)
	return wDir
}
