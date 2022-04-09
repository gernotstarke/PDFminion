package main

import (
	"pdfminion/domain"
	ui "pdfminion/gui"
)

func main() {

	domain.SetupConfiguration()
	ui.CreateMainUI_take1()

}
