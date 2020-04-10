package api

import (
	v1 "blog-api/api/v1"
	"github.com/gin-gonic/gin"
)

func NewAPI(engine *gin.Engine) {
	router := engine.Group("/api")
	{
		v1.NewV1(router)
	}
}

