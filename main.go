package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		fmt.Println("Error al abrir la base de datos:", err)
		return
	}
	defer db.Close()

	fmt.Println("Base de datos y tabla listas.")
}
