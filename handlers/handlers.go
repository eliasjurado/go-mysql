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

		var valueEmail sql.NullString
		var valuePhone sql.NullString
		err := rows.Scan(&c.Id, &c.Name, &valueEmail, &valuePhone)
		if err != nil {
			log.Fatal(err)
		}	
		
		if valueEmail.Valid{
			c.Email=valueEmail.String
		}else{
			c.Email=""
		}

		if valuePhone.Valid{
			c.Phone=valuePhone.String
		}else{
			c.Phone=""
		}



		log.Printf("%+v\n", c.Id, c.Name, c.Email, c.Phone)
	}
	log.Printf("%+v\n", "------------------")
}
