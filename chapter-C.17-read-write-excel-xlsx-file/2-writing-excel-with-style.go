package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

type M map[string]interface{}

var data = []M{
	M{"Name": "Noval", "Gender": "male", "Age": 18},
	M{"Name": "Nabila", "Gender": "female", "Age": 12},
	M{"Name": "Yasa", "Gender": "male", "Age": 11},
}

func main() {
	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Name")
	xlsx.SetCellValue(sheet1Name, "B1", "Gender")
	xlsx.SetCellValue(sheet1Name, "C1", "Age")

	err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for i, each := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["Name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["Gender"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each["Age"])
	}

	sheet2Name := "Sheet two"
	sheetIndex := xlsx.NewSheet(sheet2Name)
	xlsx.SetActiveSheet(sheetIndex)

	xlsx.SetCellValue(sheet2Name, "A1", "Hello")
	xlsx.MergeCell(sheet2Name, "A1", "B1")

	style, err := xlsx.NewStyle(`{
        "font": {
            "bold": true,
            "size": 36
        },
        "fill": {
            "type": "pattern",
            "color": ["#E0EBF5"],
            "pattern": 1
        }
    }`)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	xlsx.SetCellStyle(sheet2Name, "A1", "A1", style)

	err = xlsx.SaveAs("./file2.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
