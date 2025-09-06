package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID           int
	Username     string
	PasswordHash string
}

type PageData struct {
	Title   string
	Message string
	Users   []User // asegurar que User está definido y exportado
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Obtener todos los usuarios de la base de datos
	var users []User

	// Parsear el archivo HTML (se explicará a continuación)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		log.Printf("Error al parsear la plantilla: %v", err)
		return
	}

	data := PageData{
		Title:   "CRUD SQL LITE Y GO",
		Message: "Bienvenido a la aplicación CRUD",
		Users:   users,
	}

	// No llames a w.WriteHeader(http.StatusOK) aquí antes de Execute.
	if err := tmpl.Execute(w, data); err != nil {
		// si falla la ejecución, enviar el error (http.Error escribe la cabecera y el cuerpo)
		http.Error(w, "Error al ejecutar la plantilla: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

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

	//CREAMOS LA TABLA DE USUARIOS
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password_hash TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
		return
	}
	log.Println("Tabla de usuarios creada o ya existe")

	http.HandleFunc("/", handler)
	log.Println("Servidor escuchando en http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}

}
