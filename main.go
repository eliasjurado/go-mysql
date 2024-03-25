package main

import (
	"gomysql/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}