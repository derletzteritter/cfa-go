package ui

import (
	"cfa-go/services"
	"cfa-go/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var selectedLanguage string

func SetupUI(a fyne.App) {
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

	createResource := createResource(resourcePathInput.Text, selectedLanguage, &w)

	w.SetContent(container.New(layout.NewVBoxLayout(), header, selection, language, createResource))
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

	languages.OnChanged(selectedLanguage)

	return container.New(layout.NewVBoxLayout(), title, &languages)
}

func createResource(path, language string, w *fyne.Window) *fyne.Container {
	createButton := widget.NewButton("Create resource", func() {
		if err := utils.HasCommand("git"); err != nil {
			dialog.ShowInformation("You seem to be missing Git", "Please install Git from https://git-scm.com/", *w)
		}

		if err := utils.CloneRepo(path); err != nil {
			dialog.ShowInformation("Something went wrong", err.Error(), *w)
		}

		services.CreateTemplate(services.Template{Path: path, Language: language})
	})

	return container.New(layout.NewVBoxLayout(), createButton)
}
