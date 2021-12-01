package models

import (
	_ "database/sql"
	"strconv"
	"time"

	"github.com/go-chat/server/database"
	// "github.com/jmoiron/sqlx"

	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         uint      `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}

func FindUser(col, value string, limit int) (User, error) {
	user := User{}
	var err error
	if limit < 1 {
		err = database.DB.Get(&user, "SELECT * FROM users WHERE "+col+" = ?", value)
	} else {
		err = database.DB.Get(&user, "SELECT * FROM users WHERE "+col+" = ? LIMIT "+strconv.Itoa(limit), value)
	}

	return user, err
}
func CreateUser(name, email, password string) error {
	_, err := database.DB.Exec("INSERT INTO users (name, email, password) VALUES (?,?,?)", name, email, password)
	return err
}
func EditUserName(name string, id string) (error) {
	_, err := database.DB.Exec("UPDATE users SET name=? WHERE id=?", name, id)
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
