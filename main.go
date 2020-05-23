package main

import (
	configPkg "go-boilerplate-app/config"
	"go-boilerplate-app/endpoints"
	"go-boilerplate-app/helper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var config configPkg.Config

func main() {
	r := gin.Default()

	logConfig := configPkg.NewLogConfigFromEnvironment()
	helper.ConfigureLogger(logConfig.GetLogLevel())

	configPkg.LoadConfig(&config)

	r.GET("/ping", endpoints.PongEndpoint)

	err := r.Run()
	if err != nil {
		log.Error(err)
	}

}
