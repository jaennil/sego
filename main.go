// Package main provides various examples of Fyne API capabilities.
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"sego/data"
	"sego/left_menu"
	"sego/menu"
)

func main() {
	a := app.NewWithID("ru.dubrovskih.mmex")
	a.SetIcon(data.FyneLogo) // TODO: change
	mainWindow := a.NewWindow("Money Manager Ex")

	mainMenu := menu.Create(a, mainWindow)
	mainWindow.SetMainMenu(mainMenu)
	mainWindow.SetMaster()

	content := container.NewStack()
	setTutorial := func(t left_menu.Tab) {
		content.Objects = []fyne.CanvasObject{t.View(mainWindow)}
		content.Refresh()
	}

	split := container.NewHSplit(left_menu.Create(setTutorial, true), content)
	split.Offset = 0.2
	border := container.NewBorder(widget.NewLabel("top mainMenu TODO"), nil, nil, nil, split)
	mainWindow.SetContent(border)
	mainWindow.Resize(fyne.NewSize(640, 460))
	mainWindow.ShowAndRun()
}
