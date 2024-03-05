package main

import (
	"fmt"
	"github.com/mileusna/useragent"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("参数格式: command useragent.xlsx new .xlsx")
		os.Exit(1)
	}

	var sourceFile string = ""
	sourceFile = args[1]

	var newFile string = ""
	newFile = args[2]

	file, err := excelize.OpenFile(sourceFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for i, row := range rows {
		if len(row) == 0 {
			continue
		}

		uaStr := row[0]
		ua := useragent.Parse(uaStr)

		var device string
		if ua.Mobile {
			device = "Mobile"
		}
		if ua.Tablet {
			device = "Tablet"
		}
		if ua.Desktop {
			device = "Desktop"
		}
		if ua.Bot {
			device = "Bot"
		}

		cell := "B" + strconv.Itoa(i+1)
		err := file.SetCellValue("Sheet1", cell, device)
		if err != nil {
			panic(err)
		}
	}

	err = file.SaveAs(newFile)
	if err != nil {
		panic(err)
	}
}
