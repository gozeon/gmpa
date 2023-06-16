package utils_test

import (
	"html/template"
	"testing"

	"github.com/gozeon/gmpa/utils"
)

func TestHtml(t *testing.T) {
	htmlHelper := utils.HtmlHelper{}
	htmlHelper.SetCss(`body{margin:0;}`)
	htmlHelper.SetJs(`alert("test")`)

	tpl, err := template.New("test").Parse(`<style>{{.CssContent}}</style><script>{{.JsContent}}}}</script>`)
	if err != nil {
		t.Fatal(err)
	}
	htmlHelper.SetTemplate(tpl)

	html, err := htmlHelper.GetHtml()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(html)
}
