package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Student struct {
	Lastname  string  `csv:"last_name"`
	Firstname string  `csv:"first_name"`
	SSN       string  `csv:"ssn"`
	Test1     float32 `csv:"test_1"`
	Test2     float32 `csv:"test_2"`
	Test3     float32 `csv:"test_3"`
	Test4     float32 `csv:"test_4"`
	Final     float32 `csv:"final"`
	Grade     string  `csv:"grade"`
}

func main() {

	log.Println("Hello 123")
	var err error
	var file *os.File

	path, err := os.Getwd()
	if err != nil {
		panic("ciquan")
	}
	//log.Println(path)
	path = path + "\\cmd\\data.csv"

	// Open the CSV file
	file, err = os.Open(path)
	if err != nil {
		panic("Hi")
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		panic("ciquan")
	//	}
	//}(file)

	// Create the CSV reader
	csvReader := csv.NewReader(file)
	// read csv values using csv.Reader
	data1, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data1)
	// convert records to array of structs
	//shoppingList := createDataList(data1)

	// print the array
	//fmt.Printf("%+v\n", shoppingList)

	// Read the CSV headers
	headers, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", headers)

	// Read the CSV data rows
	var data []map[string]interface{}
	for {
		row, err := csvReader.Read()
		if err != nil {
			break
		}

		// Convert the row values to the appropriate types
		m := make(map[string]interface{})
		for i, val := range row {
			f, err := strconv.ParseFloat(val, 64)
			if err == nil {
				m[headers[i]] = f
				continue
			}

			b, err := strconv.ParseBool(val)
			if err == nil {
				m[headers[i]] = b
				continue
			}

			m[headers[i]] = val
		}

		data = append(data, m)
	}

	// print the array
	//fmt.Printf("%+v\n", data)

	// Read the CSV data
	//reader := csv.NewReader(file)
	//reader.FieldsPerRecord = -1 // Allow variable number of fields
	//data, err := reader.ReadAll()
	//if err != nil {
	//	panic(err)
	//}
	//
	//
	// Print the CSV data
	//for _, row := range data {
	//	for _, col := range row {
	//		fmt.Printf("%s,", col)
	//	}
	//	fmt.Println()
	//}

}

func createDataList(data [][]string) []Student {
	var dataList []Student
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Student
			for j, field := range line {
				//value, err := strconv.ParseFloat(field, 32)
				if j == 0 {
					rec.Lastname = field
				} else if j == 1 {
					rec.Firstname = field
				} else if j == 2 {
					rec.SSN = field
				} else if j == 3 {
					value, err := strconv.ParseFloat(field, 32)
					if err != nil {
						// do something sensible
					}
					rec.Test1 = float32(value)
				}
			}
			dataList = append(dataList, rec)
		}
	}
	return dataList
}
