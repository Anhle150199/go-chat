package jwtconfig

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"

	// "os"
	// "strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chat/server/models"
)
var jwtKey = []byte("my_secret_key")
type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}


func Create(idUser uint) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Id: idUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func Verify(w http.ResponseWriter, r *http.Request) (models.User, error) {
	c, err := r.Cookie("logged-in")
	if err != nil {
		// if err == http.ErrNoCookie {
		// 	// If the cookie is not set, return an unauthorized status
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		// // For any other type of error, return a bad request status
		// w.WriteHeader(http.StatusBadRequest)
		// return
		return models.User{}, err
	}

	// Get the JWT string from the cookie
	tknStr := c.Value
	log.Println(tknStr)
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		// if err == jwt.ErrSignatureInvalid {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		// w.WriteHeader(http.StatusBadRequest)
		return models.User{}, err
	}
	// if !tkn.Valid {
	// 	return " tkn"
	// }
	log.Println("err: ", tkn)
	idUser:= strconv.FormatUint(uint64(claims.Id), 10)

	user, err := models.FindUser("id", idUser, 1)

	if err != nil {
		return models.User{}, err
	}
	log.Println(user)
	return user, nil
}
