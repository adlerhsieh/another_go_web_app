package main

import (
	"net/http"
)

func main() {
	// http handler
	// w 是產生response用的變數
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello"))
	})

	http.ListenAndServe(":8000", nil)
}
