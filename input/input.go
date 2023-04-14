package input

import (
	"fmt"
	"strconv"
	"utilities-manager/utils"
)

func GetUserInput(question string, category string) string {
	var input string

	fmt.Println(category)
	fmt.Println("")
	fmt.Println(question)
	fmt.Println("")
	fmt.Scan(&input)

	utils.ClearScreen()

	return input
}

func GetConfirmation(address string, gas float64, water float64, electric float64) bool {
	var input string

	fmt.Println("")
	fmt.Println(address)
	fmt.Println("")
	fmt.Println("Gas      :", gas, "m3")
	fmt.Println("Water    :", water, "m3")
	fmt.Println("Electric :", electric, "kwh")
	fmt.Println("")
	fmt.Println("------------------------------")
	fmt.Println("Is this correct? (y/n)")
	fmt.Println("")
	fmt.Scan(&input)

	utils.ClearScreen()

	confirmed := input == "y"

	return confirmed
}

func Start() bool {
	var input string
	utils.ClearScreen()

	fmt.Println("")
	fmt.Println("Welcome to your utilities manager. 🏘")
	fmt.Println("")
	fmt.Println("Ready to get started? (y/n)")
	fmt.Println("")
	fmt.Scan(&input)

	confirmed := input == "y"

	if confirmed {
		utils.ClearScreen()
	}

	return confirmed
}

func End() {
	utils.ClearScreen()

	fmt.Println("")
	fmt.Println("Done! 🚀 Your data has been added to the csv.")
	fmt.Println("")
	fmt.Println("See you next month! ⏰")
	fmt.Println("")
}

func GetUserInputAsFloat(question string, category string) float64 {
	input := GetUserInput(question, category)
	value, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Sorry, '", input, "' is not a valid number value, please try again")
		fmt.Println("")
		return GetUserInputAsFloat(question, category)
	} else {
		return value
	}
}

func GetConsumedAmountsPerAptNumber(aptNumber string) [3]float64 {
	var consumedAmounts [3]float64
	questions := []string{"🔥 gas (m3):", "💧 water (m3):", "⚡ power (Kwh):"}

	for i, question := range questions {
		consumedAmounts[i] = GetUserInputAsFloat(question, aptNumber)
	}

	confirmation := GetConfirmation(aptNumber, consumedAmounts[0], consumedAmounts[1], consumedAmounts[2])

	if confirmation {
		return consumedAmounts
	} else {
		return GetConsumedAmountsPerAptNumber(aptNumber)
	}
}

func GetConsumedAmounts() []interface{} {
	var consumedAmounts []interface{}
	addresses := []string{"Street 93", "Street 93-A", "Street 93-B"}

	for _, aptNumber := range addresses {
		amounts := GetConsumedAmountsPerAptNumber(aptNumber)

		for _, amount := range amounts {
			consumedAmounts = append(consumedAmounts, amount)
		}
	}

	return consumedAmounts
}
