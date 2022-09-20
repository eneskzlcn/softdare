package web

import (
	"embed"
	coreTemplate "github.com/eneskzlcn/softdare/internal/core/html/template"
	customerror "github.com/eneskzlcn/softdare/internal/error"
	"html/template"
)

//go:embed template/include/*.gohtml template/*.gohtml

var templateFS embed.FS

var templateFuncs = template.FuncMap{
	"linkify":      coreTemplate.Linkify,
	"isLoginError": IsLoginError,
}

func ParseTemplate(name string) *template.Template {
	return coreTemplate.Parse(name, templateFuncs, templateFS,
		"template/include/*.gohtml", "template/")
}

func IsLoginError(err string) bool {
	return customerror.IsLoginErrorStr(err)
}
