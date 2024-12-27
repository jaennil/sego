package left_menu

import (
	"fyne.io/fyne/v2"
	"sego/tutorials"
)

type Tab struct {
	View func(w fyne.Window) fyne.CanvasObject
}

var (
	Tutorials = map[string]Tab{
		"Dashboard": tutorials.welcomeScreen},
		"canvas": {"Canvas",
			"See the canvas capabilities.",
			tutorials.canvasScreen,
		},
	}

	TabTree = map[string][]string{
		"Dashboard":   {"All Transactions", "Reports"},
		"Reports": {"Cash Flow"},
	}
)
