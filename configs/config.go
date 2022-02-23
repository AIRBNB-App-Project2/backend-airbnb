package configs

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Port     int
	Driver   string
	Name     string
	Address  string
	DB_Port  int
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()
	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("NAME", "crud_api_test")
	defaultConfig.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = getEnv("USERNAME", "root")
	defaultConfig.Password = getEnv("PASSWORD", "root")

	fmt.Println(defaultConfig)

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "faliq" {
		fmt.Println(value)
		return value
	}

	return fallback
}
