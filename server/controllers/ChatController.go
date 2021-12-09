package controllers

import (
	// "math/rand"
	"html/template"
	"io"
	"os"
	"strconv"

	"log"

	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chat/server/models"
	"github.com/go-chat/server/utils"
	"github.com/nouney/randomstring"
)

// Open home Go-Chat
func Index(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("User").(models.User)
	allMessage, _ := models.GetAllMessage()

	varmap := map[string]interface{}{
		"userName":   user.Name,
		"userId":     user.Id,
		"allMessage": allMessage,
	}

	tpl, _ := template.ParseGlob("public/view/*.html")
	tpl.ExecuteTemplate(w, "index.html", varmap)
}

// Edit user name
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

// Sent messages
func SendMessage(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("User").(models.User)
	r.ParseForm()
	id := r.PostFormValue("id")
	message := r.PostFormValue("message")
	id = models.Santize(id)
	message = models.Santize(message)
	err := models.InsertMessage(id, user.Name, message)
	if err != nil {
		utils.JSON(w, 400, err)
		return
	}
	utils.JSON(w, 200, "Send Succesfully")

}

// Upload file media: image and video
func UploadFile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("User").(models.User)

	log.Println("Uploading .....")
	// 32 MB is the default used by FormFile
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")

	if err != nil {
		utils.JSON(w, 400, err)
		return
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	err = os.MkdirAll("public/medias", os.ModePerm)
	if err != nil {
		utils.JSON(w, 400, err)
		return
	}

	str := randomstring.Generate(8)
	mediaFileName := str + "-" + handler.Filename
	mediafileType := handler.Header["Content-Type"][0]
	var messageType string
	if mediafileType == "image/png" || mediafileType == "image/jpg" || mediafileType == "image/gif" || mediafileType == "image/jpeg" {
		messageType = "image"
	} else if mediafileType == "video/mp4" || mediafileType == "video/mov" {
		messageType = "video"
	} else {
		utils.JSON(w, 400, "error: Type of file is Valid")
		return
	}

	dst, err := os.Create("public/medias/" + mediaFileName)
	if err != nil {
		utils.JSON(w, 400, err)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		utils.JSON(w, 400, err)
		return
	}
	err = models.UploadFile(strconv.Itoa(int(user.Id)), user.Name, mediaFileName, messageType)
	if err != nil {
		utils.JSON(w, 400, err)
		return
	}
	utils.JSON(w, 200, "Upload Succesfully")

}

// Edit content message
func EditMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Editting ....")
	r.ParseForm()
	idMessage := r.PostFormValue("idMessage")
	newMessage := r.PostFormValue("newMessage")

	idMessage = models.Santize(idMessage)
	newMessage = models.Santize(newMessage)
	log.Println("id Mess: ", idMessage)
	log.Println("new mess", newMessage)
	err := models.EditMessage(idMessage, "message_content", newMessage)
	if err != nil {
		utils.JSON(w, 400, err)
		return
	}
	utils.JSON(w, 200, "Edit Succesfully")
}

// Delete message
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting ....")
	r.ParseForm()
	idMessage := r.PostFormValue("idMessage")
	idMessage = models.Santize(idMessage)

	message, err := models.GetOneMessage(idMessage)
	if err != nil {
		utils.JSON(w, 400, err)
		return
	}

	if message.Media_file_name != "" {
		err = os.Remove("public/medias/" + message.Media_file_name)
		if err != nil {
			utils.JSON(w, 400, err)
			return
		}
	}

	err = models.DeleteMessage(idMessage)
	if err != nil {
		utils.JSON(w, 400, err)
		return
	}
	utils.JSON(w, 200, "Delete Succesfully")

}
