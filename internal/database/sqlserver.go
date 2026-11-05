package database

import (
	"database/sql"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var DB *sql.DB

func Connect() {
	connString := `
	server=172.19.28.102;
	user id=tsjweb;
	password=tsjweb;
	database=BDSistemaArchivos;
	encrypt=disable
	`

	var err error
	DB, err = sql.Open(
		"sqlserver", connString,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SQL Server conectado")
}
