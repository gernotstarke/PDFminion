package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"pdfminion/domain"
)

// Appl exposes the fyne application - mainly to enable the quit-function to stop the app.
var Appl fyne.App

// Window is the main application window
var Window fyne.Window

// CreateMainUI_take1 creates and shows the minimalistic graphical user interface.
// It creates by delegating to "Panel" functions which will create their respective panel.
func CreateMainUI_take1() {

}

// CreateMainUI_take2 creates and shows the main graphical user interface, second version.
// It creates by delegating to "Panel" functions which will create their respective panel.
func CreateMainUI_take2() {

	Appl = app.New()

	Appl.Settings().SetTheme(theme.LightTheme())
	Window = Appl.NewWindow(domain.AppName)

	content := container.New(layout.NewVBoxLayout(),
		directoriesPanel(),
		widget.NewSeparator(),
		okCancelPanel())

	Window.SetContent(content)
	Window.Resize(fyne.NewSize(600, 400))
	Window.SetFixedSize(true)
	Window.CenterOnScreen()
	Window.ShowAndRun()
}

func directoriesPanel() fyne.CanvasObject {

	dirContainer := container.New(layout.NewVBoxLayout(),
		srcDirSelectorGroup())
	dirPanel := widget.NewCard("", "Directories", dirContainer)

	return dirPanel
}

func srcDirSelectorGroup() *fyne.Container {
	srcDirField := widget.NewEntry()
	srcDirField.SetText(domain.SourceDirName())

	srcDirButton := widget.NewButton("Source", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, Window)
				return
			}
			if list == nil {
				return
			}

			_, err = list.List()
			if err != nil {
				dialog.ShowError(err, Window)
				return
			}
			srcDirField.SetText(list.Name())
			domain.SetSourceDirName(list.Name())

			if domain.CheckSrcDirectoryStatus(list.String()) {
				fmt.Printf("Folder %s :\n%s", list.Name(), list.String())
				// dialog.ShowInformation("Folder Open", out, Window)
			}
		}, Window)
	})
	srcDirButton.SetIcon(theme.FolderOpenIcon())

	domain.CheckSrcDirectoryStatus(domain.SourceDirName())
	srcDirStatusLabel := canvas.NewText("nothing selected", color.Gray{})
	srcDirStatusLabel.TextSize = 9

	inputStatus := container.New(layout.NewVBoxLayout(), srcDirField, srcDirStatusLabel)

	return container.New(layout.NewFormLayout(), srcDirButton, inputStatus)

}

func okCancelPanel() fyne.CanvasObject {

	OKButton := widget.NewButton("Process PDFs", func() {})
	OKButton.Disable()

	CancelButton := widget.NewButton("Cancel", quitApp)

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		CancelButton,
		OKButton)

	okCancelPanel := widget.NewCard("", "Processing", buttons)

	return okCancelPanel
}

func quitApp() {
	Appl.Quit()
}
