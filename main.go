package main

import (
	"bufio"
	"fmt"
	"os"
)

var Problem string
var Path string

func main() {
	Problem = "7.2"
	Path = "imports/day7_Excercise.txt"

	data := ImportFile(Path)

	switch Problem {
	case "1":
		password := Part1(data)
		fmt.Println("The password is: ", password)
	case "1.2":
		fmt.Println("The password is: ", Part2(data))
	case "2":
		fmt.Println("Sum of invalid ID's: ", Day2Part1(data))
	case "2.2":
		fmt.Println("Sum of invalid ID's: ", Day2Part2(data))
	case "3":
		fmt.Println("Total output joltage: ", Day3Part1(data))
	case "3.2":
		fmt.Println("Total output joltage: ", Day3Part2(data))
	case "4":
		fmt.Printf("The forklift can access %d rolls", Day4Part1(data))
	case "4.2":
		fmt.Printf("The forklift can now access %d rolls", Day4Part2(data))
	case "5":
		fmt.Printf("%d available ingredients are fresh.", Day5Part1(data))
	case "5.2":
		fmt.Printf("%d available ingredients are fresh.", Day5Part2(data))
	case "6":
		fmt.Printf("Grand Total: %d", Day6Part1(data))
	case "6.2":
		fmt.Printf("Grand Total: %d", Day6Part2(data))
	case "7":
		fmt.Printf("The beam is split %d times", Day7Part1(data))
	case "7.2":
		fmt.Printf("The beam is split %d times", Day7Part2(data))
	}

}

func ImportFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file at ", path)
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	return data
}
