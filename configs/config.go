package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port        int
	UserService struct {
		Address  string
		Port     string
		Endpoint string
		Url      string
	}
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
	port, errPort := strconv.Atoi(os.Getenv("APP_PORT"))
	if errPort != nil {
		log.Warn(errPort.Error())
	}

	defaultConfig.Port = port
	defaultConfig.UserService.Address = os.Getenv("BACKEND_ADDRESS")
	defaultConfig.UserService.Port = os.Getenv("ROGERDEV_BLOGSPOT_BACKEND_SERVICE_SERVICE_PORT")
	defaultConfig.UserService.Endpoint = "/users"
	defaultConfig.UserService.Url = "http://" + defaultConfig.UserService.Address + ":" + defaultConfig.UserService.Port + defaultConfig.UserService.Endpoint

	fmt.Println(defaultConfig.UserService.Url)

	return &defaultConfig
}
