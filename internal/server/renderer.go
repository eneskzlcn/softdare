package server

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

type Renderer struct {
	logger *zap.SugaredLogger
}

func NewRenderer(logger *zap.SugaredLogger) *Renderer {
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
		r.logger.Error("could not render", zap.String("template", tmpl.Name()), zap.Error(err))
		http.Error(w, fmt.Sprintf("could not render %q", tmpl.Name()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.WriteHeader(statusCode)
	_, err = buf.WriteTo(w)
	if err != nil {
		r.logger.Error("could not send", zap.String("template", tmpl.Name()), zap.Error(err))
	}
	r.logger.Infof("TEMPLATE %s RENDERED SUCCESSFULLY", tmpl.Name())
}
