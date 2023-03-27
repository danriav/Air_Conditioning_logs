package connection


import (
	"database/sql"
)


func DBConnection()(db *sql.DB){
	Driver := "mysql"
	User := "root"
	Password := ""
	Name := "aires_acondicionados"

	db, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1:3307)/"+Name)
	if err != nil {
		panic (err.Error())
	}
	return db
}
