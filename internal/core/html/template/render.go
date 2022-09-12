package template

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type Logger interface {
	Debugf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

type RenderFn func(logger Logger, w http.ResponseWriter, tmpl *template.Template, data any, statusCode int)

func Render(logger Logger, w http.ResponseWriter, tmpl *template.Template, data any, statusCode int) {
	logger.Debugf("RENDERING THE TEMPLATE %s", tmpl.Name())
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		logger.Errorf("could not render template with error: %s", err.Error())
		http.Error(w, fmt.Sprintf("could not render %q", tmpl.Name()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.WriteHeader(statusCode)
	_, err = buf.WriteTo(w)
	if err != nil {
		logger.Errorf("could not send the template with error:%s", err.Error())
		return
	}
	logger.Debugf("TEMPLATE %s RENDERED SUCCESSFULLY", tmpl.Name())
}
