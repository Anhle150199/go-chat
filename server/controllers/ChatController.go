package controllers

import (
	"html/template"
	"log"

	// "log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chat/server/models"
	"github.com/go-chat/server/utils"
	// jwtconfig "github.com/go-chat/server/jwtConfig"
)

func Index(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("User").(models.User)
	varmap := map[string]interface{}{
		"userName": user.Name,
		"userId": user.Id,
	}

	tpl, _ := template.ParseGlob("public/view/*.html")
	tpl.ExecuteTemplate(w, "index.html", varmap)
}

func EditName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.PostFormValue("id")
	name := r.PostFormValue("name")
	// validator data
	if govalidator.IsNull(name) || govalidator.IsNull(id){
		utils.JSON(w, 400, "Data can not empty")
		return
	}
	id = models.Santize(id)
	name = models.Santize(name)
	nameCheck, _ := models.FindUser("name", name, 1)
	if nameCheck.Name == name {
		utils.JSON(w, 400, "Name is exist")
		return
	}
	err:=models.EditUserName(name,id)
	if err != nil {
		utils.JSON(w, 400, "Has error")
		return 
	}
	log.Println("Edited: ", name)
	// varmap := map[string]interface{}{
	// 	"userName": name,
	// 	"userId": id,
	// }
	// tpl, _ := template.ParseGlob("public/view/*.html")
	// tpl.ExecuteTemplate(w, "index.html", varmap)
	utils.JSON(w, 200, "Edit Succesfully")

}
