package services

import "github.com/sqweek/dialog"

func OpenFolderPicker(title string) string {
	dir, err := dialog.Directory().Title(title).Browse()

	if err != nil {
		error.Error(err)
	}

	return dir
}