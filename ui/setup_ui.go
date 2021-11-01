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

		if err := utils.HasCommand("git"); err != nil {
			dialog.ShowInformation("You seem to be missing Git", "Please install Git from https://git-scm.com/", w)
		}

		if err := utils.CloneRepo(resourcePath); err != nil {
			dialog.ShowInformation("Something went wrong", err.Error(), w)
		}
	})

	projectName := widget.NewEntry()
	projectName.SetPlaceHolder("Project name")
	selection := container.New(layout.NewVBoxLayout(), resourcePathInput, openFilePicker, projectName)

	language := languageSelection()
	packages := packageSelection()

	w.SetContent(container.New(layout.NewVBoxLayout(), header, selection, language, packages))
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

func packageSelection() *fyne.Container {
	title := widget.NewLabel("Select package")

	return container.New(layout.NewVBoxLayout(), title)
}
