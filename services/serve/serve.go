package serve

import (
	"github.com/davesavic/aio/views/layout"
	"github.com/davesavic/aio/views/page"
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.GET("/", func(context *gin.Context) {
		layout.UnauthorisedLayout(
			"Login",
			page.LoginPage(),
		).Render(context, context.Writer)
	})
	return r.Run()
}
