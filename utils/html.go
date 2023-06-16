package utils

import "fmt"

type HtmlHelper struct {
	js       string
	css      string
	template string
}

func (htmlHelper *HtmlHelper) SetJs(js string) {
	htmlHelper.js = js
}

func (htmlHelper *HtmlHelper) SetCss(css string) {
	htmlHelper.css = css
}

func (htmlHelper *HtmlHelper) SetTemplate(template string) {
	htmlHelper.template = template
}

func (htmlHelper *HtmlHelper) GetHtml() string {
	return fmt.Sprintf("<style>%s</style>\n<script>%s</script>", htmlHelper.css, htmlHelper.js)
}
