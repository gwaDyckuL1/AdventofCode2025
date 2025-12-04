package main

import (
	"bufio"
	"fmt"
	"os"
)

var Problem string
var Path string

func main() {
	Problem = "3.2"
	Path = "imports/day3_Problem.txt"

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
