/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"pkg/globals"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",                // Comando que tenemos que llamar para invocar esta función
	Short: "Añadir nueva tarea", // descripción
	Long: `Esto permite añadir nuevas tareas a nuestra lista
	
	Las tareas tien los siguientes datos: nombre de la tarea, descripción e importancia`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
		// var task1 main.Task // creamos el struct

		task1.name = args[0]
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// newCmd.Flags()String()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
