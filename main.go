package main

import (
	"cfa-go/ui"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	ui.SetupUI(a)
}