package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// Appl exposes the fyne application - mainly to enable the quit-function.
	appl := app.New()

	// Window is the main application window
	window := appl.NewWindow("Take-1")

	entryField := widget.NewEntry()
	entryField.SetText("")

	statusLabel := widget.NewLabel("")

	actionButton := widget.NewButton("Action", func() {
		statusLabel.SetText(entryField.Text)
	})

	cancelButton := widget.NewButton("Cancel", func() {
		appl.Quit()
	})

	content := container.New(layout.NewVBoxLayout(),
		entryField,
		statusLabel,
		container.New(layout.NewHBoxLayout(), cancelButton, actionButton))

	window.SetContent(content)
	window.ShowAndRun()

}
