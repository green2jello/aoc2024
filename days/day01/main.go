package main

import (
	"fmt"
	"os"
	"bufio"
)

// Empty file setup for GoLang. Will remove in future just putting in structure.

func part1(input []string) int {
	// TODO: Implement part 1
	return 0
}

func part2(input []string) int {
	// TODO: Implement part 2
	return 0
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	input := readInput()
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}