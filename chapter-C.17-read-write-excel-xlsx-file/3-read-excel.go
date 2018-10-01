package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

type M map[string]interface{}

func main() {
	xlsx, err := excelize.OpenFile("./file1.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	sheet1Name := "Sheet One"

	rows := make([]M, 0)
	for i := 2; i < 5; i++ {
		row := M{
			"Name":   xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
			"Gender": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			"Age":    xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
		}
		rows = append(rows, row)
	}

	fmt.Printf("%v \n", rows)
}
