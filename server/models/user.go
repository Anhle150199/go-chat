package models

import (
	_ "database/sql"
	"strconv"
	"time"
	"html"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-chat/server/database"
)

type User struct {
	Id         uint      `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}

func FindUser(col, value string, limit int) (user User, err error) {
	if limit < 1 {
		err = database.DB.Get(&user, "SELECT * FROM users WHERE "+col+" = ?", value)
	} else {
		err = database.DB.Get(&user, "SELECT * FROM users WHERE "+col+" = ? LIMIT "+strconv.Itoa(limit), value)
	}
	return 
}
func CreateUser(name, email, password string) (err error) {
	_, err = database.DB.Exec("INSERT INTO users (name, email, password) VALUES (?,?,?)", name, email, password)
	return
}
func EditUserName(name string, id string) (err error) {
	_, err = database.DB.Exec("UPDATE users SET name=? WHERE id=?", name, id)
	return err
}
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
