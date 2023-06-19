package utils

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
)

func getAllHtml(path string) []string {
	var result []string
	err := Afs.Walk(path, func(path string, info fs.FileInfo, err error) error {
		Log.Debug(info)
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
		// fix path in windows
		// @see: https://stackoverflow.com/a/70539540
		name := filepath.Base(htmlPath)
		dir := filepath.Dir(htmlPath)

		return template.New(name).ParseFiles(getAllHtml(dir)...)
	}

	return template.New("string").Parse(`
	<head>
		<style>{{.CssContent}}</style>
	<head>
	<body>
		<script>{{.JsContent}}</script>
	</body>
	`)

}
