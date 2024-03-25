package main

import (
	"gomysql/database"
	"gomysql/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.SetFlags(0)
	db,err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	handlers.GetAllContacts(db)
	handlers.GetOneContact(db,6)
}