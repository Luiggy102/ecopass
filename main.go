package main

import (
	"os"

	"github.com/fatih/color"
)

func main() {
	calificacionesString := os.Args[1:]
	notas := comprobarArgumentos(calificacionesString)
	esVálido := comprobarNotas(notas)
	if !esVálido {
		color.Yellow("Coloque notas válidas.")
	}
}
