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
// TODO: Transformar los datos de la db a datos en un string
// TODO: Hacer una forma de que se reste la fecha actual de la fecha de vencimiento

func ShowItems(db *sql.DB) {
	// Esta funcion se encarga de mostrar las distintas tareas, hacer la tabla y todo esok

	// conseguimos la cantidad de bases de datos
	count, err := CountDB(db)

	if err != nil {
		log.Fatal(err)
	}

	if count < 1 { // Funcion para comprobar que existan almenos mas de una tarea
		//TODO: Podria hacer que envie un valor predeterminado en caso de no tener tareas

		fmt.Print("No tiene tareas")
	}

	//NOTE: Mi idea aqui es crear el tipico bucle for para enviar los datos a un slice y de ese slice a un slice 2d
	//ese slice 2d sera lo que imprimira la tabla 

	var allTask [][]string 

	for id, id >= count {
		var buffer []string
		row, err := db.Query(
			"SELECT id, name, description, priority, state, created, due", id
			).Scan(&buffer)
	}

	tableHead := []string{"ID", "Nombre", "Descripcion", "Prioridad", "Vencimiento", "Estado"} //Los datos en encabezado

	table := tw.NewWriter(os.Stdout)
	table.Header(tableHead)

}
