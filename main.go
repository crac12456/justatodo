/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"database/sql" // para usar la base de datos
	"log"          // mostrar mensajes de error
	"main.go/cmd"
	//NOTE: no te olvides de: "time"

	_ "github.com/mattn/go-sqlite3" // para tener mejor compatibilidad con sqlite
)

// TODO: mirar eso de los pquetes, me va a dar problemas tarde o temprano

func main() {
	db, err := sql.Open("sqlite3", "./todogo.db") // creamos la base de datos
	if err != nil {
		log.Fatal(err) // verificamos que no tenga ningun error en la creación
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Creamos una base de datos en sql
	_, err = db.Exec(`  
		CREATE TABLE IF NOT EXIST task (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		priority INTEGER DEFAULT 1,
		estate BOOLEAN,
		created DATETIME DEFAULT CURRENT_TIMESTAMP
		due DATE
		)
	`)

	/*
	* tenemos 4 datos,
	* name: nombre ingresado por el usuario, debe ser corto
	* description: descripción de la tarea, no es obligatoria
	* priority: prioridad de la tarea, la idea es que sea de 1 a 3
	* estate: estado, esta o no esta hecho, boleano
	* created: Fecha de creación, esta la ingresare en datasave.go
	* due: Fecha maxima, esta debe ser ingresada por el usuario
	 */

	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
