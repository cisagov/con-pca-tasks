package notifications

import (
	"bytes"
	"log"
	"text/template"
)

type Template struct {
	FirstName string
	LastName  string
}

// Render renders the template with the given text.
func (t *Template) Render(text string) string {
	var buf bytes.Buffer
	tmpl, err := template.New("template").Parse(text)
	if err != nil {
		log.Println(err.Error())
	}
	err = tmpl.Execute(&buf, t)
	if err != nil {
		log.Println(err.Error())
	}

	return buf.String()
}
