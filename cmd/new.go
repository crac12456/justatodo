/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"todogo/pkg"
)

var (
	priority    int
	description string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [task] [date]",                         // Comando que tenemos que llamar para invocar esta función
	Short: "Añadir nueva tarea, nombre y fecha en ISO", // descripción
	Long: `Esto permite añadir nuevas tareas a nuestra lista
	
	Las tareas tien los siguientes datos: nombre de la tarea, fecha de vencimiento, descripción, fecha de creación e importancia

	La fecha se ingresa en el siguiente formato: YYYY-MM-DD`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")

		var newTask pkg.Task // creamos el struct para guardar los datos
		// Datos obligatorios
		newTask.Name = args[0]
		newTask.Due = args[1]

		newTask.Created = time.ANSIC

		newTask.Description = description

		if priority > 0 || priority < 4 {
			newTask.Priority = priority
		} else {
			newTask.Priority = 1
		}

		pkg.InsertDb(newTask) // Ingresamos los datos a la base de datos
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&description, "desc", "d", "", "añade una descripcion a la tarea") // Flag para añadir una tarea
	newCmd.Flags().IntVarP(&priority, "priority", "p", 1, "prioridad de la tarea")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
