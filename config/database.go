package config

import (
	"database/sql"
	"log"
	//"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func ConnectDB(){
	db, err := sql.Open("mysql", "root:Karthika@8888@/go_products")
if err != nil {
	panic(err)
}
log.Println("Database Connected")
DB = db
}