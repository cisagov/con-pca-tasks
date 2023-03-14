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

// Render renders the template with the given string data.
func (t *Template) Render(data string) string {
	var buf bytes.Buffer
	tmpl, err := template.New("template.html").Parse(data)
	if err != nil {
		log.Println(err.Error())
	}
	err = tmpl.Execute(&buf, t)
	if err != nil {
		log.Println(err.Error())
	}

	return buf.String()
}
