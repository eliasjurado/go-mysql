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
	log.Printf("%v\n", "Lista de Contactos")
	log.Printf("%v\n", "------------------")
	for rows.Next() {
		c := models.Contact{}

		var valueEmail sql.NullString
		var valuePhone sql.NullString
		err := rows.Scan(&c.Id, &c.Name, &valueEmail, &valuePhone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			c.Email = valueEmail.String
		} else {
			c.Email = ""
		}

		if valuePhone.Valid {
			c.Phone = valuePhone.String
		} else {
			c.Phone = ""
		}
		log.Printf("ID: %v, Name: %v, Email: %v, Phone: %v\n", c.Id, c.Name, c.Email, c.Phone)
	}
	log.Printf("%v\n", "------------------")
}

func GetOneContact(db *sql.DB, id int) {
	query := "select * from contact where id = ?"
	row := db.QueryRow(query, id)

	c := models.Contact{}

	var valueEmail sql.NullString
	var valuePhone sql.NullString
	if valueEmail.Valid {
		c.Email = valueEmail.String
	} else {
		c.Email = ""
	}

	if valuePhone.Valid {
		c.Phone = valuePhone.String
	} else {
		c.Phone = ""
	}

	log.Printf("%v\n", "Lista de Contacto")
	log.Printf("%v\n", "------------------")

	err := row.Scan(&c.Id, &c.Name, &valueEmail, &valuePhone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontró ningún contacto con el ID %d", id)
		}
	}

	log.Printf("ID: %v, Name: %v, Email: %v, Phone: %v\n", c.Id, c.Name, c.Email, c.Phone)
	log.Printf("%v\n", "------------------")
}

func CreateContact(db *sql.DB, c models.Contact) {
	query := "insert into contact (name, email, phone) values(?,?,?)"

	_, err := db.Exec(query, c.Name, c.Email, c.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", "Nuevo contacto registrado con éxito")
}
