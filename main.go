package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

var (
	strNota1, strNota2, strActividades string
	nota1, nota2, actividades          int
)

func main() {
	inputN1 := huh.NewInput().
		Title("Colque la nota del primer exámen").
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
		Title("Colque la nota del segundo exámen").
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
		Title("Colque la nota de actividades").
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

	entradas := huh.NewGroup(inputN1, inputN2, inputAct).
		WithShowHelp(false).
		WithTheme(huh.ThemeBase16())

	formulario := huh.NewForm(entradas)

	if err := formulario.Run(); err != nil {
		log.Fatal(err)
	}

	nota1, _ = strconv.Atoi(strNota1)
	nota2, _ = strconv.Atoi(strNota2)
	actividades, _ = strconv.Atoi(strActividades)
}
