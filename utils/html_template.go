package utils

import (
	"html/template"
	"io/fs"
	"path"
	"strings"
)

func getAllHtml(path string) []string {
	var result []string
	err := Afs.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		return []string{}
	}
	return result
}

func GetTemplate(htmlPath string) (*template.Template, error) {
	exists, err := Afs.Exists(htmlPath)
	if err == nil && exists {
		name := path.Base(htmlPath)
		dir := path.Dir(htmlPath)
		return template.New(name).ParseFiles(getAllHtml(dir)...)
	}

	return template.New("string").Parse(`<style>{{.CssContent}}</style><script>{{.JsContent}}</script>`)

}
