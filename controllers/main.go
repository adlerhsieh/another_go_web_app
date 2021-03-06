package controllers

import (
	"bufio"
	// "fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	// Load template according to the requested path, excluding the leading "/"
	// 	requestedFile := req.URL.Path[1:]
	// 	template := templates.Lookup(requestedFile + ".html")
	//
	// 	var context interface{} = nil
	// 	context = viewmodels.Get("My Title", requestedFile)
	//
	// 	if template != nil {
	// 		template.Execute(w, context)
	// 	} else {
	// 		w.WriteHeader(404)
	// 		w.Write([]byte("404 - " + http.StatusText(404)))
	// 	}
	// })

	ic := new(indexController)
	ic.template = templates.Lookup("index.html")
	http.HandleFunc("/index", ic.get)

	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	http.HandleFunc("/categories", cc.get)

	// redirect /img/ and /css/ dir to assets/img/ and assets/css/
	http.HandleFunc("/img/", serveAssets)
	http.HandleFunc("/css/", serveAssets)
}

func serveAssets(w http.ResponseWriter, req *http.Request) {
	path := "assets" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
