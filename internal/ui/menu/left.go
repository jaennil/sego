package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	fyneTheme "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"sego/internal/domain"
	"sego/internal/ui/tabs"
	"sego/internal/ui/theme"
)

type View func(w fyne.Window, d *domain.Domain) fyne.CanvasObject

var (
	Tutorials = map[string]View{
		"Dashboard":        tabs.DashboardScreen,
		"All Transactions": tabs.AllTransactionsScreen,
		"Reports":          tabs.ReportsScreen,
	}

	TabTree = map[string][]string{
		"":          {"Dashboard"},
		"Dashboard": {"All Transactions", "Reports"},
	}
)

const preferenceCurrentTutorial = "currentTutorial"

func Left(setTutorial func(tutorial View), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return TabTree[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := TabTree[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("<BLANK>")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			_, ok := Tutorials[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(uid)
		},
		OnSelected: func(uid string) {
			if t, ok := Tutorials[uid]; ok {
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(t)
			}
		},
	}
	tree.OpenAllBranches()

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(&theme.ForcedVariant{Theme: fyneTheme.DefaultTheme(), Variant: fyneTheme.VariantDark})
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(&theme.ForcedVariant{Theme: fyneTheme.DefaultTheme(), Variant: fyneTheme.VariantLight})
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}
