package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	jwt "github.com/go-chat/server/jwtConfig"
	"github.com/go-chat/server/models"
	"github.com/go-chat/server/utils"

	"github.com/asaskevich/govalidator"
)

// Login
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl, _ := template.ParseGlob("public/view/*.html")
		tpl.ExecuteTemplate(w, "login.html", nil)
	} else {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")
		if govalidator.IsNull(email) || govalidator.IsNull(password) {
			utils.JSON(w, 400, "Data can not empty")
			return
		}
		email = models.Santize(email)
		password = models.Santize(password)

		user, err := models.FindUser("email", email, 1)

		if err != nil {
			utils.JSON(w, 401, "Email or Password incorrect")
			return
		}

		err = models.CheckPasswordHash(user.Password, password)
		if err != nil {
			utils.JSON(w, 401, "Username or Password incorrect")
			return
		}

		token, errCreate := jwt.Create(user.Id)

		if errCreate != nil {
			utils.JSON(w, 500, "Internal Server Error")
			return
		}
		// w.WriteHeader(http.StatusProxyAuthRequired)
		http.SetCookie(w, &http.Cookie{
			Name:    "logged-in",
			Value:   token,
			Path:    "/",
			HttpOnly: true,
			Secure: true,
			Expires: time.Now().Add(120 * time.Hour),
		})
		utils.JSON(w, 200, "Login Succesfully")

	}
}

// Create User
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, _ := template.ParseGlob("public/view/*.html")
		tpl.ExecuteTemplate(w, "register.html", nil)
	} else {
		r.ParseForm()
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		// validator data
		if govalidator.IsNull(name) || govalidator.IsNull(email) || govalidator.IsNull(password) {
			utils.JSON(w, 400, "Data can not empty")
			return
		}

		if !govalidator.IsEmail(email) {
			utils.JSON(w, 400, "Email is invalid")
			return
		}

		// clean data
		name = models.Santize(name)
		email = models.Santize(email)
		password = models.Santize(password)

		// check user exists
		if nameCheck, _ := models.FindUser("name", name, 1); nameCheck.Name == name {
			utils.JSON(w, 409, "User name does exists")
			return
		}
		if emailCheck, _ := models.FindUser("email", email, 1); emailCheck.Email == email {
			utils.JSON(w, 409, "Email does exists")
			return
		}

		hashedPassword, err := models.Hash(password)
		if err != nil {
			utils.JSON(w, 500, "Register has failed: PassErr")
			return
		}

		err = models.CreateUser(name, email, hashedPassword)
		if err != nil {
			utils.JSON(w, 500, "Register has failed: Create Err")
			fmt.Fprintln(w, err)
			return
		}

		utils.JSON(w, 201, "Register Succesfully")
	}
}

// Logout 
func Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("Logouting ...")
	cookie := &http.Cookie{
		Name:    "logged-in",
		Value:   "",
		Path:    "/",
		HttpOnly: true,
		Expires: time.Now().Add(0 * time.Second),
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/user/login", http.StatusMovedPermanently)
}
