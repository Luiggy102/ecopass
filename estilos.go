package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const bannerAprobado = `
$$$$$$$\                                           $$\               $$\ 
$$  __$$\                                          $$ |              $$ |
$$ |  $$ |$$$$$$\   $$$$$$$\  $$$$$$\   $$$$$$$\ $$$$$$\    $$$$$$\  $$ |
$$$$$$$  |\____$$\ $$  _____| \____$$\ $$  _____|\_$$  _|  $$  __$$\ $$ |
$$  ____/ $$$$$$$ |\$$$$$$\   $$$$$$$ |\$$$$$$\    $$ |    $$$$$$$$ |\__|
$$ |     $$  __$$ | \____$$\ $$  __$$ | \____$$\   $$ |$$\ $$   ____|    
$$ |     \$$$$$$$ |$$$$$$$  |\$$$$$$$ |$$$$$$$  |  \$$$$  |\$$$$$$$\ $$\ 
\__|      \_______|\_______/  \_______|\_______/    \____/  \_______|\__|
`

const bannerRecuperación = `
▄▄▄  ▄▄▄ . ▄▄· ▄• ▄▌ ▄▄▄·▄▄▄ .▄▄▄   ▄▄▄·  ▄▄· ▪         ▐ ▄ 
▀▄ █·▀▄.▀·▐█ ▌▪█▪██▌▐█ ▄█▀▄.▀·▀▄ █·▐█ ▀█ ▐█ ▌▪██ ▪     •█▌▐█
▐▀▀▄ ▐▀▀▪▄██ ▄▄█▌▐█▌ ██▀·▐▀▀▪▄▐▀▀▄ ▄█▀▀█ ██ ▄▄▐█· ▄█▀▄ ▐█▐▐▌
▐█•█▌▐█▄▄▌▐███▌▐█▄█▌▐█▪·•▐█▄▄▌▐█•█▌▐█ ▪▐▌▐███▌▐█▌▐█▌.▐▌██▐█▌
.▀  ▀ ▀▀▀ ·▀▀▀  ▀▀▀ .▀    ▀▀▀ .▀  ▀ ▀  ▀ ·▀▀▀ ▀▀▀ ▀█▄▀▪▀▀ █▪

`

const bannerReprobado = `

 ██▀███  ▓█████  ██▓███   ██▀███   ▒█████   ▄▄▄▄    ▄▄▄      ▓█████▄  ▒█████  
▓██ ▒ ██▒▓█   ▀ ▓██░  ██▒▓██ ▒ ██▒▒██▒  ██▒▓█████▄ ▒████▄    ▒██▀ ██▌▒██▒  ██▒
▓██ ░▄█ ▒▒███   ▓██░ ██▓▒▓██ ░▄█ ▒▒██░  ██▒▒██▒ ▄██▒██  ▀█▄  ░██   █▌▒██░  ██▒
▒██▀▀█▄  ▒▓█  ▄ ▒██▄█▓▒ ▒▒██▀▀█▄  ▒██   ██░▒██░█▀  ░██▄▄▄▄██ ░▓█▄   ▌▒██   ██░
░██▓ ▒██▒░▒████▒▒██▒ ░  ░░██▓ ▒██▒░ ████▓▒░░▓█  ▀█▓ ▓█   ▓██▒░▒████▓ ░ ████▓▒░
░ ▒▓ ░▒▓░░░ ▒░ ░▒▓▒░ ░  ░░ ▒▓ ░▒▓░░ ▒░▒░▒░ ░▒▓███▀▒ ▒▒   ▓▒█░ ▒▒▓  ▒ ░ ▒░▒░▒░ 
  ░▒ ░ ▒░ ░ ░  ░░▒ ░       ░▒ ░ ▒░  ░ ▒ ▒░ ▒░▒   ░   ▒   ▒▒ ░ ░ ▒  ▒   ░ ▒ ▒░ 
  ░░   ░    ░   ░░         ░░   ░ ░ ░ ░ ▒   ░    ░   ░   ▒    ░ ░  ░ ░ ░ ░ ▒  
   ░        ░  ░            ░         ░ ░   ░            ░  ░   ░        ░ ░  
                                                 ░            ░               
`

// https://www.patorjk.com/software/taag/#p=display&v=3&f=Bloody&t=Reprobado
// https://colorhunt.co/palettes/green

func imprimirMensaje(banner string, mensaje string, color lipgloss.Color) {
	estiloBanner := lipgloss.NewStyle().
		Foreground(color)
	estiloMensaje := lipgloss.NewStyle().
		Italic(true)
	estiloCaja := lipgloss.NewStyle().
		BorderForeground(color).
		BorderStyle(lipgloss.RoundedBorder()).
		Width(85).
		PaddingLeft(3).
		PaddingBottom(1).
		PaddingRight(3)
	fmt.Println(estiloCaja.Render(
		estiloBanner.Render(banner),
		estiloMensaje.Render("\n", mensaje),
	))
}

func tablaNotas(nota1, nota2, actividades, exámen int, mensaje string, color lipgloss.Color) string {
	t := table.New().
		Width(20).
		BorderStyle(lipgloss.NewStyle().Foreground(color))
	t.Row("Nota 1", strconv.Itoa(nota1))
	t.Row("Nota 2", strconv.Itoa(nota2))
	t.Row("Actividades", strconv.Itoa(actividades))
	t.Row("Exámen", strconv.Itoa(exámen))
	t.Row("Total", strconv.Itoa(nota1+nota2+actividades+exámen))
	return mensaje + t.String()
}
