package templ

import (
	"context"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

type Render struct {
	Code      int
	Component templ.Component
}

func (t Render) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Code)
	if t.Component != nil {
		return t.Component.Render(context.Background(), w)
	}
	return nil
}

func (t Render) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t Render) Instance(_ string, data interface{}) render.Render {
	if templData, ok := data.(templ.Component); ok {
		return &Render{
			Code:      http.StatusOK,
			Component: templData,
		}
	}
	return nil
}

func NewRenderer() render.HTMLRender {
	return Render{}
}
