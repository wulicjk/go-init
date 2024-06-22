package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"readLater-backend/infrastructure/config"
	"readLater-backend/middleware"
	"readLater-backend/router"
	"strconv"
)

var httpConf = config.Cfg.HTTPServerConf

func main() {
	baseServer := "0.0.0.0:" + strconv.Itoa(httpConf.GPort)
	log.Info("start server...\r\ngo：http://" + baseServer)

	engine := gin.Default()
	engine.Use(middleware.CrossSite)
	//记录日志
	engine.Use(middleware.NewMidLogger())
	router.InitApiRouter(engine)
	engine.Run(baseServer)
}
