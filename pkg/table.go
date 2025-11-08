package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	tw "github.com/olekukonko/tablewriter"
)

// NOTE: Esta funcion se encarga de mostrar todas las tareas
// TODO: Transformar los datos de la db a datos en un string
// TODO: Hacer una forma de que se reste la fecha actual de la fecha de vencimiento
func ShowItems(db *sql.DB) {

	// conseguimos la cantidad de bases de datos
	count, err := CountDB(db)
	if err != nil {
		log.Fatal(err)
	}

	if count < 1 {
		fmt.Print("No tiene tareas")
	}

	rows, err := db.Query("SELECT id, name, description, priority, state, created, due")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	header := []string{"ID", "Nombre", "Descripcion", "Prioridad", "Vencimiento", "Estado"}

	table := tw.NewWriter(os.Stdout)

}
