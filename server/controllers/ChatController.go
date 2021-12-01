package controllers

import (
	// "fmt"

	"html/template"
	"log"
	"net/http"

	jwtconfig "github.com/go-chat/server/jwtConfig"
)

func Index(w http.ResponseWriter, r *http.Request) {
	user, err := jwtconfig.Verify(w,r)
	if err != nil {
		log.Println("Verify fail!!")
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
	varmap := map[string]interface{}{
        "user"
    }

	tpl, _ := template.ParseGlob("front-end/view/*.html")
	tpl.ExecuteTemplate(w, "index.html", varmap)
}