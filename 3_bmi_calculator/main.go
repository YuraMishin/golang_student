package main

import (
	"fmt"
	"mishinyura.com/bmicalculator/utils"
	"strconv"
	"strings"
)

func main() {
	println("Hello, World!")

	print(utils.InputMessage1)
	weightInput, _ := reader.ReadString(utils.Delim1)
	weightInput = strings.Replace(weightInput, string(utils.Delim1), "", -1)
	weightInput = strings.Replace(weightInput, string(utils.Delim2), "", -1)
	weight, _ := strconv.ParseFloat(weightInput, 64)

	print(utils.InputMessage2)
	heightInput, _ := reader.ReadString(utils.Delim1)
	heightInput = strings.Replace(heightInput, string(utils.Delim1), "", -1)
	heightInput = strings.Replace(heightInput, string(utils.Delim2), "", -1)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI: %.2f", bmi)
}
