package main

import (
	"bufio"
	"fmt"
	"gomysql/database"
	"gomysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.SetFlags(0)
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	//Create tables
	database.CreateTable(db, models.UserSchema)

	defer db.Close()

	// for {
	// 	fmt.Println("\nMenú:")
	// 	fmt.Println("1. Listar contactos")
	// 	fmt.Println("2. Obtener contacto por ID")
	// 	fmt.Println("3. Crear nuevo contacto")
	// 	fmt.Println("4. Actualizar contacto")
	// 	fmt.Println("5. Eliminar contacto")
	// 	fmt.Println("6. Salir")
	// 	fmt.Print("Seleccione una opción: ")

	// 	// Leer la opción seleccionada por el usuario
	// 	var option int
	// 	fmt.Scanln(&option)

	// 	// Ejecutar la opción seleccionada
	// 	switch option {
	// 	case 1:
	// 		handlers.GetAllContacts(db)
	// 	case 2:
	// 		fmt.Print("Ingrese el ID del contacto: ")
	// 		var idContact int
	// 		fmt.Scanln(&idContact)
	// 		handlers.GetOneContact(db, idContact)
	// 	case 3:
	// 		newContact := inputContactDetails()
	// 		handlers.CreateContact(db, newContact)
	// 	case 4:
	// 		updatedContact := inputContactDetails()
	// 		handlers.UpdateContact(db, updatedContact)
	// 	case 5:
	// 		fmt.Print("Ingrese el ID del contacto que quiere eliminar: ")
	// 		var idContact int
	// 		fmt.Scanln(&idContact)
	// 		handlers.DeleteContact(db, idContact)
	// 	case 6:
	// 		fmt.Println("Saliendo del programa...")
	// 		return
	// 	default:
	// 		fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
	// 	}
	// }
}

// Función para ingresar los detalles del contacto desde la entrada estándar
func inputContactDetails() models.Contact {
	// Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo electrónico del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el número de teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
