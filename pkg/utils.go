package pkg

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: falta añadir una forma de mostrar los datos
// TODO: ahora falta rescribir esta funcion para que funcione con el struct

func InsertDb(name, description, date string, priority int) { // Introducimos datos a la base de datos
	db, err := sql.Open("sqlite3", "./todogo.db") // creamos la base de datos

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if len(date) != 10 { // verificadno que la fecha este ingresada correctamente, esto no funcina muy bien, pero es algo
		log.Fatal("No se ha ingresado la fecha correctamente")
	}

	if len(name) < 1 { // El usuario tiene que ingresar si o si un nombre
		log.Fatal("No se ha ingresado un nombre")
		return
	}
	if priority > 3 || priority < 0 { // El error lo dice todo, en la base de datos el minimo es 1, pero igual
		log.Fatal("se ha ingresado incorrectamente la prioridad")
	}

	// Ingresamos los datos a la base de datos (dah)
	_, err = db.Exec(`
		INSERT INTO task (name, description, priority, due_date)
		VALUES (?, ?, ?, ?)
		`, name, description, priority, date)

	if err != nil {
		log.Fatal(err)
	}

}

// NOTE: Esto es la base de los datos, estos son los datos que tendra cada tarea
type Task struct {
	Name        string
	Description string
	Priority    int
	Estate      bool
	Created     string
	Due         string
}

// NOTE: Esta funcion se encarga de mostrar todas las tareas
func ShowItems() {

}
