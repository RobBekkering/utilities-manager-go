package main

import (
	input "utilities-manager/input"
	storage "utilities-manager/storage"
	utils "utilities-manager/utils"
)

func main() {
	var row []interface{}

	start := input.Start()

	if !start {
		return
	}

	row = append(row, utils.GetDate())
	row = append(row, input.GetConsumedAmounts()...)

	stringRow := utils.GetStringRow(row)
	storage.AppendRowToFile(stringRow)

	input.End()
}
