package tabs

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
)

func DashboardScreen(_ fyne.Window) fyne.CanvasObject {
    return widget.NewLabel("Dashboard")
}
