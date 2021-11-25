package controllers

import (
	// "fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	CheckCookie(w, r)
	tmp, _ := template.ParseFiles("front-end/view/index.html")
	tmp.Execute(w, nil)
}
func CheckCookie(w http.ResponseWriter, r *http.Request)  {
	_, err := r.Cookie("logged-in")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}

}