package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GivenInput struct {
	Instruction string
	Value       int
}

func main() {
	const filename = "./input/input.txt"

	data := FileReader(filename)
	formatedData := FormatData(data)

	result := Compute(formatedData)
	fmt.Println("The result of first part is: ", result)

	resultPartTwo := ComputePartTwo(formatedData)
	fmt.Println("The result of second part is: ", resultPartTwo)
}

func Compute(data []GivenInput) int {
	horizontal := 0
	depth := 0
	for _, given := range data {
		if given.Instruction == "forward" {
			horizontal += given.Value
		}
		if given.Instruction == "up" {
			depth -= given.Value
		}
		if given.Instruction == "down" {
			depth += given.Value
		}
	}
	return horizontal * depth
}

func ComputePartTwo(data []GivenInput) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, given := range data {
		if given.Instruction == "forward" {
			horizontal += given.Value
			depth += aim * given.Value
		}
		if given.Instruction == "up" {
			aim -= given.Value
		}
		if given.Instruction == "down" {
			aim += given.Value
		}
	}
	return horizontal * depth
}

// error handling functino
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads input data from txt
func FileReader(filename string) string {
	dat, err := os.ReadFile(filename)
	check(err)
	data := string(dat)

	return data
}

// formats data to work better
func FormatData(data string) []GivenInput {
	inputLines := strings.Split(data, "\n")

	myInputs := []GivenInput{}
	for _, inputLine := range inputLines {
		splittedLine := strings.Split(string(inputLine), " ")

		instruction := splittedLine[0]
		value, err := strconv.Atoi(splittedLine[1])
		check(err)

		newInput := &GivenInput{Instruction: instruction, Value: value}
		myInputs = append(myInputs, *newInput)
	}
	return myInputs
}
