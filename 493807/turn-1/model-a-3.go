package main

import (
	"fmt"
	"html/template"
)

func main() {
	t := template.Must(template.New("greeting").Parse(`Hello, {{ .Name }}! Your score is {{ .Score }}.`))
	data := struct {
		Name  string
		Score int
	}{
		Name:  "Alice",
		Score: 95,
	}
	var message bytes.Buffer
	t.Execute(&message, data)
	fmt.Println(message.String())
}
