package web

import (
	"bytes"
	"embed"
	"fmt"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

//go:embed template/include/*.gohtml template/*.gohtml
var templateFS embed.FS

func parseTemplate(name string) *template.Template {
	tmpl := template.New(name)
	//first parses the layout the files that all template should include
	tmpl = template.Must(template.ParseFS(templateFS, "template/include/*.gohtml"))
	//parsing the exact template and merge with layout then return.
	return template.Must(tmpl.ParseFS(templateFS, "template/"+name))
}

func (h *Handler) renderTemplate(w http.ResponseWriter, tmpl *template.Template, data any, statusCode int) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		h.logger.Error("could not render", zap.String("template", tmpl.Name()), zap.Error(err))
		http.Error(w, fmt.Sprintf("could not render %q", tmpl.Name()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.WriteHeader(statusCode)
	_, err = buf.WriteTo(w)
	if err != nil {
		h.logger.Error("could not send", zap.String("template", tmpl.Name()), zap.Error(err))
	}
}
