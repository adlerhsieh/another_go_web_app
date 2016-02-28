package controllers

import (
	"github.com/adlerhsieh/web/viewmodels"
	"net/http"
	"text/template"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.Get("Categoreis Page", "categories")
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vm)
}
