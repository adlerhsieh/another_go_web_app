package main

import (
	"net/http"
	"os"
	"text/template"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    // Load template according to the requested path
		requestedFile := req.URL.Path[1:]
		template := templates.Lookup(requestedFile + ".html")
		if template != nil {
			template.Execute(w, nil)
		} else {
			w.WriteHeader(404)
      w.Write([]byte("404 - " + http.StatusText(404)))
		}
	})
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	basePath := "templates"
  // open and close dir
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	// -1 means every file
	templatePathsRaw, _ := templateFolder.Readdir(-1)

  // read every file and append filepath to a slice
  templatePaths := new([]string)
	for _, pathInfo := range   templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	// ... at the end will brings in all values in a slice as args
	result.ParseFiles(*templatePaths...)
	return result
}
