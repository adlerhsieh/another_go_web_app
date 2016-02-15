package main

import (
	// "fmt"
	"net/http"
	"os"
	"text/template"
  "github.com/adlerhsieh/web/controllers"
)

func main() {
	templates := populateTemplates()
	// Move http.HandleFunc() to controllers
	controllers.Register(templates)
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
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	// ... at the end will brings in all values in a slice as args
	result.ParseFiles(*templatePaths...)
	return result
}
