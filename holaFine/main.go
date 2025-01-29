package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/internal/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hola Fyne")

	myWindow.SetContent(widget.NewLabel("Hola Mundo"))

	myWindow.ShowAndRun()

}
