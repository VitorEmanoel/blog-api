package post

import (
	"blog-api/crud"
	"github.com/gin-gonic/gin"
)

type PostsContext struct {

}

func NewPostCrud() crud.RestCrud {
	return &PostsContext{}
}

func NewPost(group *gin.RouterGroup) {
	var postsCrud = NewPostCrud()
	router := group.Group("/posts")
	{
		crud.RegisterCrud(postsCrud, router)
	}
}

func (p *PostsContext) Find(ctx *gin.Context) {
}

func (p *PostsContext) FindAll(ctx *gin.Context) {
}

func (p *PostsContext) Create(ctx *gin.Context) {
}

func (p *PostsContext) Delete(ctx *gin.Context) {
}

func (p *PostsContext) Update(ctx *gin.Context) {
}

func (p *PostsContext) Patch(ctx *gin.Context) {
}
