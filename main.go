package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {
	// 1ro con argumentos
	calificaciones := os.Args[1:]
	// comprobar la cantidad de argumentos
	if len(calificaciones) < 2 || len(calificaciones) > 3 {
		log.Fatal(color.RedString("\t¡Cantidad de argumentos invalida!"))
	}
}
