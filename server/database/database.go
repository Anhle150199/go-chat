package database

import (
	_ "database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func ConnectDB() (error) {
        var err error
        db, err = sqlx.Open("mysql", "root:@tcp(mysql)/go-chat?charset=utf8mb4&parseTime=true&loc=Local&time_zone=%27Asia%2FHo_Chi_Minh%27")
        if err != nil {
                log.Fatal("sql error: ", err)
        }

        for {
                err = db.Ping()
                if err != nil {
                        time.Sleep(1 * time.Second)
                        log.Println("watting Mysql ....")
                } else {
                        break
                }
        }
        fmt.Println("Database Connected!")
        return err
	}