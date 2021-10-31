package main

import (
	"cfa-go/services"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("CFA")
	w.Resize(fyne.Size{Width: 1200, Height: 720})
	w.SetFixedSize(true)

	title := widget.NewLabelWithStyle("Resource Creation Wizard", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	header := container.New(layout.NewVBoxLayout(), title, layout.NewSpacer())

	resourcePathInput := widget.NewEntry()
	openFilePicker := widget.NewButton("Browse folder", func() {
		resourcePath := services.OpenFolderPicker("Select resource path")

		resourcePathInput.SetText(resourcePath)
	})

	projectName := widget.NewEntry()
	projectName.SetPlaceHolder("Project name")

	selection := container.New(layout.NewVBoxLayout(), resourcePathInput, openFilePicker, projectName)
	language := languageSelection()

	w.SetContent(container.New(layout.NewVBoxLayout(), header, selection, language))
	w.ShowAndRun()
}

func languageSelection() *fyne.Container {
	title := widget.NewLabel("Select langauge")

	languages := widget.RadioGroup{
		Horizontal: false,
		Required:   true,
		Options:    []string{"Lua", "JavaScript", "TypeScript"},
		Selected:   "Lua",
	}

	return container.New(layout.NewVBoxLayout(), title, &languages)
}
