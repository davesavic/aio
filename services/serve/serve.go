package serve

import (
	templEngine "github.com/davesavic/aio/services/templ"
	view "github.com/davesavic/aio/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() error {
	r := gin.New()
	r.HTMLRender = templEngine.NewRenderer()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "landing", view.LandingPage())
	})

	return r.Run()
}
