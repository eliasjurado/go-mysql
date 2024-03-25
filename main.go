package main

import (
	"gomysql/database"
	"gomysql/handlers"
	"gomysql/models"
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
	contact := models.Contact{
		Name: "Hugo Ascencios",
		Email: "vigo@gmail.com",
		Phone: "+51987453723",
	}
	handlers.CreateContact(db,contact)
	handlers.GetAllContacts(db)
	handlers.GetOneContact(db,6)
}