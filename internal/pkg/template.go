package pkg

import (
	"embed"
	"github.com/mvdan/xurls"
	"html/template"
)

//go:embed template/include/*.gohtml template/*.gohtml

var templateFS embed.FS

var templateFuncs = template.FuncMap{
	"linkify": linkify,
}

func ParseTemplate(name string) *template.Template {
	tmpl := template.New(name).Funcs(templateFuncs)
	tmpl = template.Must(tmpl.ParseFS(templateFS, "template/include/*.gohtml"))
	return template.Must(tmpl.ParseFS(templateFS, "template/"+name))
}

func linkify(s string) template.HTML {
	s = template.HTMLEscapeString(s)
	return template.HTML(xurls.Relaxed.
		ReplaceAllString(s, `<a href ="$0" target="_blank" rel="noopener noreferror">$0</a>`))
}
