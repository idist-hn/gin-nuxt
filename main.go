package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/martinlindhe/notify"
	"idist-core/app/providers/configProvider"
	"idist-core/app/providers/jobsProvider"
	"idist-core/app/providers/loggerProvider"
	"idist-core/app/providers/mongoProvider"
	"idist-core/app/providers/redisProvider"
	"idist-core/app/providers/routerProvider"
	"idist-core/app/providers/socketProvider"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		loggerProvider.GetLogger().Info("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	gin.DisableConsoleColor()
	if *environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	configProvider.Init(*environment)
	loggerProvider.Init()
	routerProvider.Init(r)
	mongoProvider.Init()
	redisProvider.Init()
	jobsProvider.Init()
	socketProvider.Init()

	defer mongoProvider.CloseMongoDB()
	///* Run server */

	loggerProvider.GetLogger().Info("Version: " + configProvider.GetConfig().GetString("app.server.version"))
	loggerProvider.GetLogger().Info("Branch: " + configProvider.GetConfig().GetString("app.server.branch"))
	buildBatch := "Version: " + configProvider.GetConfig().GetString("app.server.version") + "\n" + "Branch: " + configProvider.GetConfig().GetString("app.server.branch")

	notify.Notify("Portal API CMS", "Building success", buildBatch, "")
	if err := r.Run(fmt.Sprintf("%s:%s",
		configProvider.GetConfig().GetString("app.server.host"),
		configProvider.GetConfig().GetString("app.server.port"))); err != nil {
		loggerProvider.GetLogger().Error(err.Error())
	}

	notify.Notify("Start Frontend", "Building success", buildBatch, "")
}
