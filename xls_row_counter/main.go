package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//countRows()

	path := "path/to/file.xlsx"
	err := collectProductIDs(path)
	if err != nil {
		log.Fatal(err)
	}

}

func countRows() {
	files, err := ioutil.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}

	total := 0

	for _, file := range files {
		path := fmt.Sprintf("data/%s", file.Name())

		f, err := excelize.OpenFile(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer func() {
			// Close the spreadsheet.
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		sheetName := f.GetSheetName(0)
		rows, err := f.GetRows(sheetName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(file.Name(), len(rows))

		total = total + len(rows)
	}
	fmt.Println(total)
}

func collectProductIDs(path string) error {
	xls, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := xls.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	output, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

	sheetName := xls.GetSheetName(0)
	rows, err := xls.Rows(sheetName)
	if err != nil {
		return err
	}

	for rows.Next() {
		cells, err := rows.Columns()
		if err != nil {
			return err
		}

		id := fmt.Sprintf("%q,", cells[0])
		_, err = output.WriteString(id + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
