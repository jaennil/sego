package tabs

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"reflect"
	"sego/internal/domain"
	"sego/internal/entity"
	"time"
)

func AllTransactionsScreen(w fyne.Window, d *domain.Domain) fyne.CanvasObject {
	transactions, err := d.GetTransactions()
	if err != nil {
		dialog.ShowError(err, w)
		return widget.NewLabel("error")
	}
	data := convertTransactionsTo2DStrings(transactions)

	return &widget.Table{
		Length: func() (rows int, cols int) {
			return len(data), len(data[0])
		},
		CreateCell: func() fyne.CanvasObject {
			return widget.NewLabel(data[0][0])
		},
		UpdateCell: func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
		CreateHeader: func() fyne.CanvasObject {
			return widget.NewLabel("Amount")
		},
		UpdateHeader: func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText("Amount")
		},
	}
}

func convertTransactionsTo2DStrings(transactions []entity.Transaction) [][]string {
	var result [][]string

	for _, transaction := range transactions {
		val := reflect.ValueOf(transaction)
		var fieldValues []string

		// Loop through each field in the struct
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)

			// Convert field value to string
			var valueStr string
			switch field.Kind() {
			case reflect.String:
				valueStr = field.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueStr = fmt.Sprintf("%d", field.Int())
			case reflect.Float32, reflect.Float64:
				valueStr = fmt.Sprintf("%f", field.Float())
			case reflect.Bool:
				valueStr = fmt.Sprintf("%t", field.Bool())
			case reflect.Struct:
				if field.Type() == reflect.TypeOf(time.Time{}) {
					valueStr = field.Interface().(time.Time).Format(time.RFC3339)
				}
			default:
				valueStr = fmt.Sprintf("%v", field.Interface())
			}

			fieldValues = append(fieldValues, valueStr)
		}

		// Append the field values of the current transaction to the result
		result = append(result, fieldValues)
	}

	return result
}
