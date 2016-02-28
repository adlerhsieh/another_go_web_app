package controllers

import (
	"github.com/adlerhsieh/web/viewmodels"
	"net/http"
	"text/template"
)

type indexController struct {
	template *template.Template
}

func (this *indexController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.Get("Index Page", "index")
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vm)
}
