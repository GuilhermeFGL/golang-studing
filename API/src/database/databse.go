package database

import (
	"database/sql"
	"example.com/m/v2/API/src/configuration"
	_ "github.com/go-sql-driver/mysql" // implicit import
	"log"
)

func Connect() (*sql.DB, error) {
	log.Println("Connecting to database")

	db, err := sql.Open("mysql", configuration.DbStringConnection)

	if err != nil {
		log.Fatal("Unable to connect to MySQL:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Unable to ping MySQL:", err)
		return nil, err
	}

	return db, nil
}
