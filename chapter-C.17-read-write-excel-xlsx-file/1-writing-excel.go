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

	err = xlsx.SaveAs("./file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
