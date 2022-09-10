package oops

import (
	"github.com/eneskzlcn/softdare/internal/pkg"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

const DomainName = "oops"

type ErrData struct {
	Err error
}
type Renderer interface {
	RenderTemplate(w http.ResponseWriter, template *template.Template, data any, statusCode int)
}

func RenderPage(renderer Renderer, logger *zap.SugaredLogger, w http.ResponseWriter, data ErrData, statusCode int) {
	logger.Debugf("Rendering the ooops page.")
	tmpl, err := pkg.ParseTemplate(DomainName)
	if err != nil {
		logger.Error("can not parse oops template", zap.Error(err))
		return
	}
	renderer.RenderTemplate(w, tmpl, data, statusCode)
}
