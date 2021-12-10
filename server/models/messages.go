package models

import (
	_ "database/sql"
	"time"

	"github.com/go-chat/server/database"
)

type Messages struct {
	Id              int       `json:"id" db:"id"`
	User_id         int       `json:"user_id" db:"user_id"`
	User_name       string    `json:"user_name" db:"user_name"`
	Message_type    string    `json:"message_type" db:"message_type"`
	Message_content string    `json:"message_content" db:"message_content"`
	Media_file_name string    `json:"media_file_name" db:"media_file_name"`
	Created_at      time.Time `json:"created_at" db:"created_at"`
	Updated_at      time.Time `json:"updated_at" db:"updated_at"`
}

func GetAllMessage() (message []Messages, err error) {
	err = database.DB.Select(&message, "SELECT * FROM messages")
	return
}

func GetOneMessage(id string) (message Messages, err error) {
	err = database.DB.Get(&message, "SELECT * FROM messages WHERE id = ?", id)
	return 
}

func InsertMessage(userId, userName, message string) (err error) {
	_, err = database.DB.Exec("INSERT INTO messages (user_id, user_name, message_type, message_content, media_file_name) VALUES (?,?,\"message\", ?, \"\")", userId, userName, message)
	return 
}

func UploadFile(userId, userName, fileName, fileType string) (err error) {
	_, err = database.DB.Exec("INSERT INTO messages (user_id, user_name, message_type, message_content, media_file_name) VALUES (?,?,?,\"\", ?)", userId, userName, fileType, fileName)
	return
}

func EditUserNameMessage(id, col, value string) (err error) {
	_, err = database.DB.Exec("UPDATE messages SET "+col+"=? WHERE user_id=?", value, id)
	return
}

func EditMessage(id, col, value string) (err error) {
	_, err = database.DB.Exec("UPDATE messages SET "+col+"=? WHERE id=?", value, id)
	return 
}

func DeleteMessage(id string) (err error) {
	_, err = database.DB.Exec("DELETE FROM messages WHERE id=?", id)
	return 
}
