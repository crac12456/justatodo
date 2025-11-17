package pkg

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: encontrar una forma de comprovar la fecha

func InsertDb(task Task) {
	// Introducimos los datos a la db

	db, err := sql.Open("sqlite3", "./todogo.db") // creamos la base de datos

	if err != nil { // Comprobamos errores en la creacion de la base de datos
		log.Fatal(err)
	}
	defer db.Close()

	if len(task.Created) != 10 { // verificamos que la fecha este ingresada correctamente, esto no funcina muy bien, pero es algo
		// WARN: deberia cambiar esto lo mas pronto posible
		log.Fatal("No se ha ingresado la fecha correctamente")
	}

	if len(task.Name) < 1 { // El usuario tiene que ingresar si o si un nombre
		log.Fatal("No se ha ingresado un nombre")
		return
	}
	if task.Priority > 3 || task.Priority < 0 { // El error lo dice todo, en la base de datos el minimo es 1, pero igual
		log.Fatal("se ha ingresado incorrectamente la prioridad")
	}

	// Ingresamos los datos
	_, err = db.Exec(`
		INSERT INTO task (name, due, description, priority, created)
		VALUES (?, ?, ?, ?, ?)
		`, task.Name, task.Due, task.Description, task.Priority, task.Created)

	if err != nil {
		log.Fatal(err)
	}

}

// NOTE: Esto es la base de los datos, estos son los datos que tendra cada tarea
type Task struct {
	Id          int
	Name        string
	Description string
	Priority    int
	State       bool
	Created     string
	Due         string
}

// NOTE: Esta funcion consigue la cantidad de tablas en la db

func CountDB(db *sql.DB) (int, error) {
	/*
	 * Esta funcion es utilizada con el fin de hacer comprovaciones en otras partes del codigo,
	 * Devuelve la cantidad de tareas
	 */

	var count int    //el numero de tablas
	var query string // La peticion a la base de datos

	// La peticion a la db para obtener la cantidad de registros
	query = `
		SELECT COUNT(*)
		FROM sqlite_master
		WHERE type = 'table'
		AND name NOT LIKE 'sqlite_%'
	`

	err := db.QueryRow(query).Scan(&count) // hacemos la peticion a la base de datos

	if err != nil { // Si no hay tablas, regresamos 0 o err
		return 0, err
	}

	return count, nil
}
