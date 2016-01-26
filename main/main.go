package main

import (
	"io/ioutil"
	"net/http"
)

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// path := "public/" + req.URL.Path
	// data, err := ioutil.ReadFile(string(path))
	data, err := ioutil.ReadFile("./public" + req.URL.Path)
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })

	// The following custom handler is to replace the above basic HandleFunc
	http.Handle("/", new(MyHandler))

	// nil means no multiplexing
	http.ListenAndServe(":8000", nil)
}
