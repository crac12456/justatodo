package pkg

import (
	//"container/heap"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	tw "github.com/olekukonko/tablewriter" // Tablewriter, permite crear tablas de forma sencilla en ASCII
)

// NOTE: Esta funcion se encarga de mostrar todas las tareas
// TODO: Hacer una forma de que se reste la fecha actual de la fecha de vencimiento

func ShowItems(db *sql.DB) {
	// Esta funcion se encarga de mostrar las distintas tareas en una tabla
	// utilizamos la libreria tablewriter

	// conseguimos la cantidad de bases de datos
	count, err := CountDB(db)

	if err != nil {
		log.Fatal(err)
	}

	if count < 1 { // Funcion para comprobar que existan almenos mas de una tarea
		//TODO: Podria hacer que envie un valor predeterminado en caso de no tener tareas

		fmt.Print("No tiene tareas")
	}

	allTask := []Task{}

	rows, err := db.Query("SELECt * FROM task") //Seleccionamos las tareas
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

		// Declaramos variables con los datos que conseguiremos de la db
		var name, description, created, due string
		var id, priority int
		var state bool

		// Hacemos la peticion de los datos
		err := rows.Scan(&id, &name, &description, &priority, &state, &created, &due)
		if err != nil {
			log.Fatal(err)
		}

		// Esta variable tendra los datos temporalmente
		buffer := Task{Id: id, Name: name, Description: description,
			Priority: priority, State: state,
			Created: created, Due: due}

		allTask = append(allTask, buffer) // Añadimos el buffer como nueva tarea al struct
	}

	defer rows.Close()

	tableHead := []string{"ID", "Nombre", "Descripcion", "Prioridad", "Vencimiento", "Estado"} //Los datos en encabezado

	table := tw.NewWriter(os.Stdout)
	table.Header(tableHead)

}
