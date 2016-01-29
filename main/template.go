package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content Type", "text/html")
		templates := template.New("template")
		templates.New("doc").Parse(doc)
		templates.New("header").Parse(header)
		context := Context{
			"Google",
			"John",
			[3]string{"Apple", "Lemon", "Orange"},
		}
		templates.Lookup("doc").Execute(w, context)
	})
	http.ListenAndServe(":8000", nil)
}

// For all template logic see https://golang.org/pkg/text/template/
const doc = `
<!DOCTYPE html>
<html>
{{template "header" .Title}}
	<body>
	  <h1>Hi, {{.Name}}</h1>
		<h3>The fruits are:</h3>
		<ul>
			{{range .Fruit}}
			  <li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>
`

const header = `
  <head>
	  {{if eq . "Google"}}
			<title>This is Google</title>
		{{else}}
			<title>A {{.}}</title>
		{{end}}
	</head>
`

type Context struct {
	Title string
	Name  string
	Fruit [3]string
}
