package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func CsvParser_WithInbuilt() {

	data := `id,name,city,role
101,Aarav Sharma,Faridabad,Analyst
102,Priya Nair,Bengaluru,Designer
103,"Rohan Verma, Jr.",Mumbai,Engineer
104,Ananya Gupta,Delhi,Manager`

	csvParser := csv.NewReader(strings.NewReader(data))
	rows, err := csvParser.ReadAll()
	if err != nil {
		fmt.Printf("Found Error in CSV format")
	}

	for i, row := range rows {
		fmt.Printf("row no is = %d and value is = %v\n", i, row)
	}

}
