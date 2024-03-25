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
		Id: "6",
		Name: "Hugo Ormeño",
		Email: "vigodelperu@gmail.com",
		Phone: "+51987453666",
	}
	handlers.UpdateContact(db,contact)
	handlers.GetAllContacts(db)
	handlers.GetOneContact(db,6)
	handlers.DeleteContact(db,6)
}