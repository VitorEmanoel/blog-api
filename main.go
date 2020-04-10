package main

import (
	"blog-api/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	server := gin.Default()
	api.NewAPI(server)
	err := server.Run()
	if err != nil {
		logrus.Error("Error in start HTTP Server")
	}
}
