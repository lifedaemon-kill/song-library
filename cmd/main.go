package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"song-library/configs"
	"song-library/db"
	"song-library/logger"
)

var Log *logrus.Logger //

func main() {
	var err error
	Log, err = logger.GetNewLogger()

	if err != nil {
		panic(err)
	}
	Log.Info("Init logger successful")

	db, err := db.NewDB(configs.GetConfig())
	if err != nil {
		Log.Fatal("Init database was failed", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		Log.Fatal("Init database was failed", err)
	}
	Log.Info("Init database successful")

	router := gin.Default()
	Log.Info("Init server successful")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong!")
	})

}
