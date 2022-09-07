package pkg

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed template/include/*.gohtml template/*.gohtml

var templateFS embed.FS

func ParseTemplate(name string) (*template.Template, error) {
	tmpl := template.New(name)
	//first parses the layout the files that all template should include
	tmpl, err := template.ParseFS(templateFS, "template/include/*.gohtml")
	if err != nil {
		fmt.Printf("Error occurred when parsing the including templates with name %s\n", name)
		return nil, err
	}
	//parsing the exact template and merge with layout then return.
	tmpl, err = tmpl.ParseFS(templateFS, "template/"+name+".gohtml")
	if err != nil {
		fmt.Printf("Error occurred when parsing the exact template with name %s\n", name)
		return nil, err
	}
	return tmpl, nil
}
