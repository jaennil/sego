package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"sego/internal/domain"
	"sego/internal/ui/data"
	"sego/internal/ui/menu"
)

type UI struct {
	domain *domain.Domain
}

func NewUI(domain *domain.Domain) *UI {
	return &UI{domain: domain}
}

func (u *UI) Start() {
	a := app.NewWithID("ru.dubrovskih.mmex")
	a.SetIcon(data.FyneLogo) // TODO: change
	mainWindow := a.NewWindow("Money Manager Ex")

	mainMenu := menu.Main(a, mainWindow)
	mainWindow.SetMainMenu(mainMenu)
	mainWindow.SetMaster()

	content := container.NewStack()
	setTutorial := func(t menu.View) {
		content.Objects = []fyne.CanvasObject{t(mainWindow, u.domain)}
		content.Refresh()
	}

	topMenu := menu.TopMenu(mainWindow, u.domain)

	split := container.NewHSplit(menu.Left(setTutorial, true), content)
	split.Offset = 0.2
	border := container.NewBorder(topMenu, nil, nil, nil, split)
	mainWindow.SetContent(border)
	mainWindow.Resize(fyne.NewSize(640, 460))
	mainWindow.ShowAndRun()
}
