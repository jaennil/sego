package menu

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"sego/internal/domain"
	"sego/internal/entity"
	"sego/internal/ui/custom_widgets"
	"strconv"
	"time"
)

func TopMenu(w fyne.Window, d *domain.Domain) fyne.CanvasObject {
	newTransactionButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		transaction := &entity.Transaction{}

		dateInput := dateInput(w)

		status := widget.NewSelect(entity.Statuses(), func(s string) {
			status, err := entity.NewStatus(s)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			transaction.Status = status
		})
		status.SetSelectedIndex(0)

		typ := widget.NewSelect(entity.Types(), func(s string) {
			typ, err := entity.NewType(s)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			transaction.Typ = typ
		})
		typ.SetSelectedIndex(0)

		amount := widget.NewEntry()

		items := []*widget.FormItem{
			widget.NewFormItem("Date", dateInput),
			widget.NewFormItem("Status", status),
			widget.NewFormItem("Type", typ),
			widget.NewFormItem("Amount", amount),
		}
		dialog.ShowForm("New Transaction", "Confirm", "Cancel", items, func(b bool) {
			fmt.Println(dateInput.Text)
			timestamp, err := time.Parse("2006/01/02", dateInput.Text)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			transaction.Date = timestamp

			f, err := strconv.ParseFloat(amount.Text, 64)
			transaction.Amount = f

			err = d.CreateTransaction(transaction)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
		}, w)
	})

	return container.NewHBox(
		newTransactionButton,
	)
}

func dateInput(w fyne.Window) *widget.Entry {
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

	return dateInput
}
