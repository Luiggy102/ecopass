package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	strNota1, strNota2, strActividades, strExámen string
	nota1, nota2, actividades, exámen             int
)

func main() {
	presentación := huh.NewNote().
		Title("\n----- Ecopass -----").
		Description("TUI calcular si pasas o no")

	inputN1 := huh.NewInput().
		Title("Colque la nota del 1ER EXÁMEN").
		Description("nota entre 0 a 20").
		Placeholder("18").
		Value(&strNota1).
		Validate(func(s string) error {
			intTransformado, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("Coloque un número en la nota 1")
			}
			if intTransformado < 0 || intTransformado > 20 {
				return errors.New("Coloque un número entre 0 a 20 en la nota 1")
			}
			return nil
		})

	inputN2 := huh.NewInput().
		Title("Colque la nota del 2DO EXÁMEN").
		Description("nota entre 0 a 20").
		Placeholder("16").
		Value(&strNota2).
		Validate(func(s string) error {
			intTransformado, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("Coloque un número en la nota 2")
			}
			if intTransformado < 0 || intTransformado > 20 {
				return errors.New("Coloque un número entre 0 a 20 en la nota 2")
			}
			return nil
		})

	inputAct := huh.NewInput().
		Title("Colque la nota de ACTIVIDADES").
		Description("nota entre 0 a 30").
		Placeholder("22").
		Value(&strActividades).
		Validate(func(s string) error {
			intTransformado, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("Coloque un número en nota de actividades")
			}
			if intTransformado < 0 || intTransformado > 30 {
				return errors.New("Coloque un número entre 0 a 30 en la nota de actividades")
			}
			return nil
		})

	inputExam := huh.NewInput().
		Title("Colque la nota del EXÁMEN FINAL").
		Description("nota entre 0 a 30").
		Placeholder("20").
		Value(&strExámen).
		Validate(func(s string) error {
			intTransformado, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("Coloque un número en nota del exámen")
			}
			if intTransformado < 0 || intTransformado > 30 {
				return errors.New("Coloque un número entre 0 a 30 en la nota de exámen")
			}
			return nil
		})

	grupo1 := huh.NewGroup(presentación,
		inputN1,
		inputN2,
		huh.NewNote().Next(true).NextLabel("Siguente"),
	).WithShowHelp(false)

	grupo2 := huh.NewGroup(huh.NewNote(),
		inputAct,
		inputExam,
		huh.NewNote().Next(true).NextLabel("Finalizar"),
	).WithShowHelp(false)

	formulario := huh.NewForm(grupo1, grupo2)

	if err := formulario.Run(); err != nil {
		log.Fatal(err)
	}

	nota1, _ = strconv.Atoi(strNota1)
	nota2, _ = strconv.Atoi(strNota2)
	actividades, _ = strconv.Atoi(strActividades)
	exámen, _ = strconv.Atoi(strExámen)
	total := nota1 + nota2 + actividades + exámen

	switch {
	case total >= 70 && actividades >= 18:
		mensaje := fmt.Sprintf("¡Felicidades pasaste con %d!\n", total)
		colorTabla := lipgloss.Color("#88D66C")
		imprimirMensaje(
			bannerAprobado,
			tablaNotas(nota1, nota2, actividades, exámen, mensaje, colorTabla),
			colorTabla,
		)
	case total >= 70 && actividades < 18:
		mensaje := fmt.Sprintf("Sacaste %d en total\nTe falto puntos en actividades para pasar directo.\n", total)
		colorTabla := lipgloss.Color("#F6FB7A")
		imprimirMensaje(
			bannerRecuperación,
			tablaNotas(nota1, nota2, actividades, exámen, mensaje, colorTabla),
			colorTabla,
		)
	case (total > 60 && total < 70) && actividades >= 18:
		mensaje := fmt.Sprintf("Sacaste %d en total\nTe falto %d puntos para pasar directo.\n", total, 70-total)
		colorTabla := lipgloss.Color("#F6FB7A")
		imprimirMensaje(
			bannerRecuperación,
			tablaNotas(nota1, nota2, actividades, exámen, mensaje, colorTabla),
			colorTabla,
		)
	case total <= 60:
		mensaje := fmt.Sprintf("Te quedaste con %d.\nTe faltaron %d puntos para pasar.\n", total, 70-total)
		colorTabla := lipgloss.Color("#EE4E4E")
		imprimirMensaje(
			bannerReprobado,
			tablaNotas(nota1, nota2, actividades, exámen, mensaje, colorTabla),
			colorTabla,
		)
	}

}
