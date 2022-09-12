package server

import (
	"bytes"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"html/template"
	"net/http"
)

type Renderer struct {
	logger logger.Logger
}

func NewRenderer(logger logger.Logger) *Renderer {
	if logger == nil {
		return nil
	}
	return &Renderer{logger: logger}
}

func (r *Renderer) RenderTemplate(w http.ResponseWriter, tmpl *template.Template, data any, statusCode int) {
	r.logger.Infof("RENDERING THE TEMPLATE %s", tmpl.Name())
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		r.logger.Errorf("could not render template with error: %s", err.Error())
		http.Error(w, fmt.Sprintf("could not render %q", tmpl.Name()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.WriteHeader(statusCode)
	_, err = buf.WriteTo(w)
	if err != nil {
		r.logger.Errorf("could not send the template with error:%s", err.Error())
		return
	}
	r.logger.Infof("TEMPLATE %s RENDERED SUCCESSFULLY", tmpl.Name())
}
