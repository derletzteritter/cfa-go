package ui

import (
	"cfa-go/network"

	"fyne.io/x/fyne/widget"
)

func RenderPackages(options []string) {
	entry := widget.NewCompletionEntry(options)

	entry.OnChanged = func(s string) {
		result := network.GetPackages(s)

		entry.SetOptions(result.Results[0])
	}
}
