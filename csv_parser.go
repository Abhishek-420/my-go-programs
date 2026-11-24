package main

import "fmt"

func CSVParser() {

	data := `id,name,city,role
101,Aarav Sharma,Faridabad,Analyst
102,Priya Nair,Bengaluru,Designer
103,"Rohan Verma, Jr.",Mumbai,Engineer
104,Ananya Gupta,Delhi,Manager`

	var rows [][]string
	var field []byte
	var row []string
	inquotes := false
	for i := 0; i < len(data); i++ {
		ch := data[i]

		if ch == '"' {
			if inquotes && i+1 < len(data) && data[i+1] == '"' {
				field = append(field, '"')
				i++
			} else {
				inquotes = !inquotes
			}
		} else if ch == ',' && !inquotes {
			row = append(row, string(field))
			field = field[:0]
		} else if ch == '\n' && !inquotes {
			row = append(row, string(field))
			field = field[:0]
			rows = append(rows, row)
			row = nil
		} else if ch == '\r' && !inquotes {

		} else {
			field = append(field, ch)
		}
	}

	if len(field) > 0 && len(row) > 0 {
		row = append(row, string(field))
		rows = append(rows, row)

	}

	for i, row := range rows {
		fmt.Printf("row no is = %d and value is = %v \n", i, row)
	}

}
