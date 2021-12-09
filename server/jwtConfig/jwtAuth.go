package jwtconfig

import (
	"log"
	"net/http"
	"strconv"
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
	expirationTime := time.Now().Add(120 * time.Hour)
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
	for _, cookie := range r.Cookies() {
        log.Println( cookie.Name)
    }

	// log.Println(c)
	if err != nil {
		log.Println("no cookie")
		return models.User{}, err
	}

	// Get the JWT string from the cookie
	tknStr := c.Value
	// log.Println(tknStr)
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return models.User{}, err
	}
	log.Println("tkn: ", tkn)
	idUser:= strconv.FormatUint(uint64(claims.Id), 10)

	user, err := models.FindUser("id", idUser, 1)

	if err != nil {
		return models.User{}, err
	}
	// log.Println(user)
	return user, nil
}
