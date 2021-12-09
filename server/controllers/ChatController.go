package controllers

import (
	"html/template"
	// "io/ioutil"
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
	allMessage, _ := models.GetAllMessage()
	
	varmap := map[string]interface{}{
		"userName": user.Name,
		"userId":   user.Id,
		"allMessage": allMessage,
	}

	tpl, _ := template.ParseGlob("public/view/*.html")
	tpl.ExecuteTemplate(w, "index.html", varmap)
}

func EditName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.PostFormValue("id")
	name := r.PostFormValue("name")
	// validator data
	if govalidator.IsNull(name) || govalidator.IsNull(id) {
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
	err := models.EditUserName(name, id)
	if err != nil {
		utils.JSON(w, 400, "Has error")
		return
	}
	err = models.EditUserNameMessage(id, "user_name", name)
	if err != nil {
		utils.JSON(w, 400, "Has error")
		return
	}
	log.Println("Edited: ", name)
	utils.JSON(w, 200, "Edit Succesfully")
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("User").(models.User)
	r.ParseForm()
	id := r.PostFormValue("id")
	message := r.PostFormValue("message")
	id = models.Santize(id)
	message = models.Santize(message)
	err:=models.InsertMessage(id,user.Name, message)
	if err != nil {
		utils.JSON(w, 400, err)
		return 
	}
	utils.JSON(w, 200, "Send Succesfully")

}

func SendFileMedia(w http.ResponseWriter, r *http.Request)  {
	
}

func EditMessage(w http.ResponseWriter, r *http.Request)  {
	log.Println("Editting ....")
	r.ParseForm()
	idMessage := r.PostFormValue("idMessage")
	newMessage := r.PostFormValue("newMessage")

	idMessage = models.Santize(idMessage)
	newMessage = models.Santize(newMessage)
	log.Println("id Mess: ",idMessage)
	log.Println("new mess", newMessage)
	err := models.EditMessage(idMessage, "message_content", newMessage)
	if err != nil {
		utils.JSON(w, 400, err)
		return 
	}
	utils.JSON(w, 200, "Edit Succesfully")
}
func DeleteMessage(w http.ResponseWriter, r *http.Request)  {
	log.Println("Editting ....")
	r.ParseForm()
	idMessage := r.PostFormValue("idMessage")
	idMessage = models.Santize(idMessage)

	err := models.DeleteMessage(idMessage)
	if err != nil {
		utils.JSON(w, 400, err)
		return 
	}
	utils.JSON(w, 200, "Delete Succesfully")

}
