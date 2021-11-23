package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "xxx")
	tmp, _ := template.ParseFiles("front-end/index.html")
	tmp.Execute(w, nil)
}