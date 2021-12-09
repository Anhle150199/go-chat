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
	Message_content string    `json:"message_content" db:"message_content"`
	Media_file_name string    `json:"media_file_name" db:"media_file_name"`
	Stamp_id        string    `json:"stamp_id" db:"stamp_id"`
	Created_at      time.Time `json:"created_at" db:"created_at"`
	Updated_at      time.Time `json:"updated_at" db:"updated_at"`
}

func GetAllMessage() ([]Messages, error) {
	message := []Messages{}
	err := database.DB.Select(&message, "SELECT * FROM messages")
	return message, err
}

func InsertMessage(userId, userName, message string) error {
	_, err := database.DB.Exec("INSERT INTO messages (user_id, user_name, message_content, media_file_name, stamp_id) VALUES (?,?, ?, \"\", \"\")", userId, userName, message)
	return err
}

func EditUserNameMessage(id, col, value string) error {
	_, err := database.DB.Exec("UPDATE messages SET "+col+"=? WHERE user_id=?", value, id)
	return err
}

func EditMessage(id, col, value string) error {
	_, err := database.DB.Exec("UPDATE messages SET "+col+"=? WHERE id=?", value, id)
	return err
}

func DeleteMessage(id string) error {
	_, err := database.DB.Exec("DELETE FROM messages WHERE id=?", id)
	return err
}
