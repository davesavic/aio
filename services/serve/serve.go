package serve

import (
	templEngine "github.com/davesavic/aio/services/templ"
	"github.com/davesavic/aio/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() error {
	r := gin.New()
	r.HTMLRender = templEngine.NewRenderer()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", view.AuthorisedLayout())
	})

	return r.Run()
}
