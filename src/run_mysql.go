package main
import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func Connect(conn_str string)  *sqlx.DB {
	db, err := sqlx.Connect("mysql", conn_str)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
func Execute(db *sqlx.DB, sql string) (err error){
	_, err = db.Exec(sql)
	return
}
func Close(db *sqlx.DB) {
	db.Close();
}