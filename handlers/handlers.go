package handlers

import (
	"database/sql"
	"gomysql/models"
	"log"
)

func GetAllContacts(db *sql.DB) {
	query := "select * from contact"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	log.Printf("%+v\n", "Lista de Contactos")
	log.Printf("%+v\n", "------------------")
	for rows.Next() {
		c := models.Contact{}
		err := rows.Scan(&c.Id, &c.Name, &c.Email, &c.Phone)
		if err != nil {
			log.Fatal(err)
		}		
		log.Printf("%+v\n", c.Id, c.Name, c.Email, c.Phone)
	}
	log.Printf("%+v\n", "------------------")
}
