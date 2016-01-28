package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if err == nil {
			tmpl.Execute(w, nil)
		}
	})
	http.ListenAndServe(":8000", nil)
}

const doc = `
<!DOCTYPE html>
<html>
  <head>
	  <title>A Title</title>
	</head>
	<body>
	  <h1>Hi</h1>
	</body>
</html>
`
