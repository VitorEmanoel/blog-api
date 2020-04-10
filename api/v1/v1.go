package v1

import (
	"blog-api/api/v1/post"
	"github.com/gin-gonic/gin"
)

func NewV1(group *gin.RouterGroup) {
	route := group.Group("/v1")
	{
		post.NewPost(route)
	}
}
