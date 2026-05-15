package excel

import (
	"fmt"

	maintenancev1 "go-kozadayev-exercise/task10/api/proto"

	"github.com/xuri/excelize/v2"
)

func ExportWorkOrders(orders []*maintenancev1.WorkOrder) (string, error) {
	f := excelize.NewFile()
	defer f.Close()

	sheet := "WorkOrders"
	f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	headers := []string{"ID", "Car ID", "Description", "Total Cost"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, o := range orders {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), o.Id)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), o.CarId)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), o.Description)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), o.TotalCost)
	}

	fileName := "work_orders.xlsx"
	if err := f.SaveAs(fileName); err != nil {
		return "", err
	}
	return fileName, nil
}
