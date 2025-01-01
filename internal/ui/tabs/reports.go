package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"sego/internal/domain"
)

func ReportsScreen(_ fyne.Window, d *domain.Domain) fyne.CanvasObject {
	return widget.NewLabel("Reports")
}
