package crud

import (
	"github.com/gin-gonic/gin"
)

type RestCrud interface {
	Find(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Patch(ctx *gin.Context)
}

func RegisterCrud(crud RestCrud, group *gin.RouterGroup) {
	group.GET("", crud.FindAll)
	group.GET("/:id", crud.Find)
	group.POST("", crud.Create)
	group.DELETE("/:id", crud.Delete)
	group.PUT("/:id", crud.Update)
	group.PATCH("/:id", crud.Patch)
}
