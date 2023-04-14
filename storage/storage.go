package storage

import (
	"encoding/csv"
	"fmt"
	"os"
)

func AppendRowToFile(row []string) {
	file, openError := os.OpenFile("utilities.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if openError != nil {
		fmt.Println(openError)
	}

	writer := csv.NewWriter(file)

	writer.Write(row[:])
	writer.Flush()
}
