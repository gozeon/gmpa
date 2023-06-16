package utils

import (
	"bytes"
	"html/template"
)

type HtmlHelper struct {
	js       string
	css      string
	template *template.Template
}

func (htmlHelper *HtmlHelper) SetJs(js string) {
	htmlHelper.js = js
}

func (htmlHelper *HtmlHelper) SetCss(css string) {
	htmlHelper.css = css
}

func (htmlHelper *HtmlHelper) SetTemplate(template *template.Template) {
	htmlHelper.template = template
}

func (htmlHelper *HtmlHelper) GetHtml() (string, error) {
	var tpl bytes.Buffer
	data := struct {
		JsContent  template.JS
		CssContent template.CSS
	}{
		CssContent: template.CSS(htmlHelper.css),
		JsContent:  template.JS(htmlHelper.js),
	}
	err := htmlHelper.template.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
