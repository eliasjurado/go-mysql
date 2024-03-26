package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dns := fmt.Sprintf("%v:%v@(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	//Abrir conexión a la db
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	//Verificar conexión a la db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Printf("%+v\n", "conexion exitosa")
	return db, nil
}

func CreateTable(db *sql.DB, schema string)  {
	db.Exec(schema)
}