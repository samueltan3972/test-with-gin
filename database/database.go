package database

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var err error
//var datasource = "root:@tcp(127.0.0.1:3306)/test"
var datasource = "test:abcd123_@tcp(test-demo-mysql.mysql.database.azure.com:3306)/test"
var db *gorp.DbMap

//Init ...
func Init() {
	db, err = ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
}

//ConnectDB ...
func ConnectDB() (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", datasource)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	return dbmap, nil

}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
