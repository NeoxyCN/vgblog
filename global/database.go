package global

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(db *sqlx.DB) {
	db_url := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_URL + ")/" + DB_DB
	database, err := sqlx.Open("mysql", db_url)

	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}
