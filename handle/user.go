package handle

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DB_TABLE = "user_basic"
)

// 此处参数首字母必须大写，公有
type user_basic struct {
	Id        int    `db:"id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	Nickname  string `db:"nickname"`
	Usergroup string `db:"usergroup"`
	Avatar    string `db:"avatar"`
	Status    string `db:"status"`
}

// TODO: use generic for parameter
func SelectUserByValueAndKey(db *sqlx.DB, value string, key string) (result []user_basic, err error) {
	var users []user_basic
	//sql :=
	//	"SELECT id,username,password,nickname,usergroup,avatar FROM user_basic WHERE  id=?"

	sql := "SELECT " + key + " FROM " + DB_TABLE + " WHERE " + key + "=?"
	err = db.Select(&users, sql, value)
	if err != nil {
		fmt.Println("select user value failed", err)
		return nil, err
	}
	// TODO: return的时候没有error也要return吗
	return users, err
}

func SelectMax() {

}

func Delete() {

}

func Insert() {

}
