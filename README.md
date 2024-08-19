# ecopass

Breve TUI que use con el fin de poner en practicar el
framework de Go: ![huh?](https://pkg.go.dev/github.com/charmbracelet/huh@v0.5.2) junto con ![lipgloss](https://pkg.go.dev/github.com/charmbracelet/lipgloss@v0.12.1).  

Realizado un formulario que pide notas a un estudiante
y dependiendo de ciertas con condiciones da como salida tres tipo de mensajes.  

- Aprobado
- Recuperación
- Reprobado

Así como también muestra una tabla con las notas ingresadas.

## Instalación
La instalación se puede hacer por medio de Go.
```bash
go install github.com/Luiggy102/ecopass@latest
```

## Demostración
![alt text](./media/demostración.gif) 

### Mensajes de salida
Aprobado
![alt text](./media/pasaste.png) 
En recuperación
![alt text](./media/recuperación.png) 
Reprobado
![alt text](./media/reprobado.png) 
