package main

import (
	"bufio"
	"fmt"
	"os"
)

var Problem string
var Path string

func main() {
	Problem = "1.2"
	Path = "imports/day1_Problem.txt"

	data := ImportFile(Path)

	switch Problem {
	case "1":
		password := Part1(data)
		fmt.Println("The password is: ", password)
	case "1.2":
		fmt.Println("The password is: ", Part2(data))
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
