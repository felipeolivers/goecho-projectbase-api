package config

import (
	"os"
)

type Server struct {
	AppName        string
	AppVersion     string
	AppPortServer  string
	AppEnvironment string
}

var Env = Server{}

func init() {
	loadEnvironments()
	
	Env.AppName = os.Getenv("APP_NAME")
	Env.AppVersion = os.Getenv("APP_VERSION")
	Env.AppPortServer = os.Getenv("APP_PORTSERVER")
	Env.AppEnvironment = os.Getenv("APP_ENVIRONMENT")
}
