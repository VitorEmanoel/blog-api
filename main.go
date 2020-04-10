package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	server := gin.Default()
	err := server.Run()
	if err != nil {
		logrus.Error("Error in start HTTP Server")
	}
}
