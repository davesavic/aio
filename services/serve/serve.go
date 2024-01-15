package serve

import (
	"fmt"
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

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login", view.Login())
	})
	r.POST("/login", func(ctx *gin.Context) {
		fmt.Println(ctx.PostForm("email"))
		fmt.Println(ctx.PostForm("password"))

		ctx.Redirect(http.StatusFound, "/")
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", view.Dashboard())
	})

	return r.Run()
}
