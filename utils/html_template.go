package utils

import (
	"html/template"
	"path"
)

func GetTemplate(htmlPath string) (*template.Template, error) {
	exists, err := Afs.Exists(htmlPath)
	if err == nil && exists {
		name := path.Base(htmlPath)
		return template.New(name).ParseFiles(htmlPath)
	}

	return template.New("string").Parse(`<style>{{.CssContent}}</style><script>{{.JsContent}}</script>`)

}
