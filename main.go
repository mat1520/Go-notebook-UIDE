package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//CREAMOS LA CONEXION A LA BASE DE DATOS MEDIANTE EL ARCHIVO DE USER
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//VERIFICAR SU CONECCION  A LA DB
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
