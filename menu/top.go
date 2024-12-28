package menu

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    custom_widget "sego/custom_widgets"
    "time"
)

func Top(w fyne.Window) fyne.CanvasObject {
    newTransactionButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
        dateInput := widget.NewEntry()
        dateInput.SetPlaceHolder("0000/00/00")
        dateInput.ActionItem = widget.NewButtonWithIcon("", theme.MoreHorizontalIcon(), func() {
            when := time.Now()

            if dateInput.Text != "" {
                t, err := time.Parse("2006/01/02", dateInput.Text)
                if err == nil {
                    when = t
                }
            }

            datepicker := custom_widget.NewDatePicker(when, time.Monday, func(when time.Time, ok bool) {
                if ok {
                    dateInput.SetText(when.Format("2006/01/02"))
                }
            })

            dialog.ShowCustomConfirm(
                "Choose date",
                "Ok",
                "Cancel",
                datepicker,
                datepicker.OnActioned,
                w,
            )
        })

        status := widget.NewSelect([]string{"Unreconciled", "Void"}, func(s string) {
            fmt.Println(s)
        })
        status.SetSelectedIndex(0)

        typ := widget.NewSelect([]string{"Withdrawal", "Deposit"}, func(s string) {
            fmt.Println(s)
        })
        typ.SetSelectedIndex(0)

        amount := widget.NewEntry()

        items := []*widget.FormItem{
            widget.NewFormItem("Date", dateInput),
            widget.NewFormItem("Status", status),
            widget.NewFormItem("Type", typ),
            widget.NewFormItem("Amount", amount),
        }
        dialog.ShowForm("New Transaction", "Confirm", "Cancel", items, func(b bool) {}, w)
    })

    return container.NewHBox(
        newTransactionButton,
    )
}
