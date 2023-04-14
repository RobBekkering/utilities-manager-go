package utils

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func GetStringRow(arr []interface{}) []string {
	var newList []string

	for _, x := range arr {
		newList = append(newList, fmt.Sprint(x))
	}

	return newList
}

func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func GetDate() string {
	time := time.Now()
	date := fmt.Sprint(time.Day(), "-", int(time.Month()), "-", time.Year())

	return date
}
