package www

import (
	"github.com/monkeydioude/moon/template"
)

// GetHome display home
func GetHome() ([]byte, int, error) {
	tmp := template.NewEngine("www/layout.html")

	tmp.BindValues(template.Javascript("/public/morab.js"))

	tmp.
		BindTemplate("www/home.html", "Content").
		BindValues(&template.HTML{"Title": "Rumor Spreader"})

	tmp.BindValues(template.Javascript("/public/test.js"))

	return tmp.Render(), 200, nil
}
