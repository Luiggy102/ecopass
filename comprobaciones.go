package main

import (
	"log"
	"strconv"

	"github.com/fatih/color"
)

func comprobarArgumentos(calificacionesString []string) (calificaciones []int) {
	// comprobar la cantidad de argumentos
	if len(calificacionesString) < 2 || len(calificacionesString) > 3 {
		color.Yellow("¡Cantidad de argumentos inválida!")
	}
	// comprobar si se coloco los números (tranformar)
	for _, calificacionString := range calificacionesString {
		calificacionInt, err := strconv.Atoi(calificacionString)
		if err != nil {
			log.Fatal(color.RedString("Error al parsear el argumento dado.\nColoque un argumento número."))
		}
		calificaciones = append(calificaciones, calificacionInt)
	}
	return
}

func comprobarNotas(calificaciones []int) (sonLasNotasVálidas bool) {
	sonLasNotasVálidas = true
	switch len(calificaciones) {
	case 2: // 1er y 2do Parcial, saber lo mínimo para actividades y exámenes
		// chequear notas validad
		if calificaciones[0] < 0 || calificaciones[0] > 20 {
			sonLasNotasVálidas = false
		}
		if calificaciones[1] < 0 || calificaciones[1] > 20 {
			sonLasNotasVálidas = false
		}
	case 3: // 1er, 2do y Actividades, saber lo mínimo para el exámen
		if calificaciones[0] < 0 || calificaciones[0] > 20 {
			sonLasNotasVálidas = false
		}
		if calificaciones[1] < 0 || calificaciones[1] > 20 {
			sonLasNotasVálidas = false
		}
		if calificaciones[2] < 0 || calificaciones[2] > 30 {
			sonLasNotasVálidas = false
		}
	}
	return
}
