package main

import (
	"langgo/api"
	"langgo/app"
	"langgo/app/middleware"
	"langgo/bootstrap"
	"langgo/bootstrap/plugins"
)

//	@title			业务框架LangGo接口
//	@version		1.0
//	@description	LangGo相关接口
//	@contact.name	qinguoyi
//	@contact.email	qinguoyiwork@gmail.com
//	@host			127.0.0.1:8890
//	@BasePath		/
func main() {
	// config log
	bootstrap.NewConfig("conf/config.yaml")
	lgLogger := bootstrap.NewLogger()

	// plugins DB Redis Minio
	plugins.NewPlugins()
	defer plugins.ClosePlugins()

	// middleware
	middleware.Init()

	// router
	api.InitCoreRouter()
	server := app.NewHttpServer()

	// app run-app
	application := app.NewApp(lgLogger.Logger, server)
	application.RunServer()
}
