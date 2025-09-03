package main

import (
	"database/sql"
	"fmt"
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
	log.Println("Conectado a la base de datos")

	//AGREGAR INFORMACION A LA TABLA USUARIOS
	//* _, err = db.Exec("INSERT INTO users (id, username, password_hash) VALUES (?, ?, ?)", 1, "Mat", "381020")
	//*if err != nil {
	//*	log.Fatal(err)
	//*}

	//CONSULTA  A LA TABLA USUARIOS
	rows, err := db.Query("SELECT id, username, password_hash FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, passwordHash string
		if err := rows.Scan(&id, &username, &passwordHash); err != nil {
			log.Fatal(err)
		}
		log.Printf("User: %d, %s, %s\n", id, username, passwordHash)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	// Insertar mediantelas varibale definidas
	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", username, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User inserted successfully")
}
