package database

import (
	_ "database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB()  {
	var err error
	DB, err = sqlx.Open("mysql", "root:@tcp(mysql)/go-chat?parseTime=true&loc=Local&time_zone=%27Asia%2FHo_Chi_Minh%27")
	if err != nil {
		log.Fatal("sql error: ", err)
	}

	for {
		err = DB.Ping()
		if err != nil {
			time.Sleep(1 * time.Second)
			log.Println("watting Mysql ....")
		} else {
			break
		}
	}
	fmt.Println("Database Connected!")
	
}