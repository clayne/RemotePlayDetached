package fynegui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/smaTc/RemotePlayDetached/executor"
)

// AppItem struct
type AppItem struct {
	App executor.App
	widget.Box
}

// NewAppItem func
func NewAppItem(app executor.App) fyne.Widget {
	item := &AppItem{App: app}
	item.ExtendBaseWidget(item)
	item.Box = *widget.NewHBox()
	item.Box.Append(layout.NewSpacer())

	argsContent := app.Args

	if argsContent == "" {
		argsContent = "none"
	}

	argsLabel := widget.NewLabel("Args: " + argsContent)
	item.Box.Append(argsLabel)
	item.Box.Append(layout.NewSpacer())

	runButton := widget.NewButton("Run", func() {
		fmt.Println("Run App")
		err := executor.RunApp(app)
		if err != nil {
			TextPopup(err.Error(), "Error:")
		}
	})
	item.Box.Append(runButton)

	editButton := widget.NewButton("Edit", func() {
		fmt.Println("Edit App")
		editApp(app)
	})
	item.Box.Append(editButton)

	deleteButton := widget.NewButton("Delete", func() {
		fmt.Println("Delete App")
		executor.DeleteApp(app)
		refreshMainWindow()
	})
	item.Box.Append(deleteButton)

	return item
}

func editApp(oldApp executor.App) {
	editWindow := rpd.NewWindow("Edit App")
	editWindow.Resize(fyne.NewSize(400, 150))

	nameEntry := NewButtonEntry()
	nameEntry.SetText(oldApp.Name)

	pathEntry := NewButtonEntry()
	pathEntry.SetText(oldApp.GamePath)

	argsEntry := NewButtonEntry()
	argsEntry.SetText(oldApp.Args)

	protonEntry := NewButtonEntry()
	protonEntry.SetText(oldApp.ProtonPath)

	prefixEntry := NewButtonEntry()
	prefixEntry.SetText(oldApp.WinePrefixPath)

	compatEntry := NewButtonEntry()
	compatEntry.SetText(oldApp.CompatDataPath)

	name := widget.NewFormItem("Name", nameEntry)
	path := widget.NewFormItem("Game Path", pathEntry)
	args := widget.NewFormItem("Args", argsEntry)
	proton := widget.NewFormItem("Proton Path", protonEntry)
	prefix := widget.NewFormItem("Prefix Path", prefixEntry)
	compat := widget.NewFormItem("Compat Path", compatEntry)
	form := widget.NewForm(name, path, args, proton, prefix, compat)

	cancelButton := widget.NewButton("Cancel", func() {
		editWindow.Close()
	})

	okButton := widget.NewButton("OK", func() {
		appName := nameEntry.Text
		appPath := pathEntry.Text
		argsString := argsEntry.Text
		protonPath := protonEntry.Text
		prefixPath := prefixEntry.Text
		compatPath := compatEntry.Text

		if appName == "" || appPath == "" {
			return
		}

		newApp := executor.App{Name: appName, GamePath: appPath, Args: argsString, ProtonPath: protonPath, WinePrefixPath: prefixPath, CompatDataPath: compatPath}
		editWindow.Close()
		executor.EditApp(oldApp, newApp)
		refreshMainWindow()
	})

	fileExlporerButton := widget.NewButton("File Explorer", func() {
		FileExplorer(pathEntry)
	})

	nameEntry.SetConfirmButton(okButton)
	pathEntry.SetConfirmButton(okButton)
	argsEntry.SetConfirmButton(okButton)
	protonEntry.SetConfirmButton(okButton)
	prefixEntry.SetConfirmButton(okButton)
	compatEntry.SetConfirmButton(okButton)

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), okButton, layout.NewSpacer(), fileExlporerButton, layout.NewSpacer(), cancelButton)

	editWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), form, layout.NewSpacer(), buttons))

	editWindow.Show()
}
