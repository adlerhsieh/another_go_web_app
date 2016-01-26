package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "./public" + req.URL.Path
	// The following line reads the requested file directly, which takes huge memory usage
	// Instead, we use buffer to prevent it. os simply opens the file and read it only when budio sends it to a reader
	// data, err := ioutil.ReadFile(path)
	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)
		var ContentType string
		if strings.HasSuffix(path, ".css") {
			ContentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			ContentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			ContentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			ContentType = "image/png"
		} else {
			ContentType = "text/plain"
		}
		w.Header().Add("Content Type", ContentType)
		bufferedReader.WriteTo(w)
		// w.Write(data)
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
