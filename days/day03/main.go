package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	testdata := false
	printdebug := false
	var data string
	if testdata {
		data = readFile("days/day03/test2_input.txt")
	} else {
		data = readFile("days/day03/input.txt")
	}
	// print first 5 chars
	println(data[:5])
	matches := regexMatch(data, `mul\([0-9]{1,3},[0-9]{1,3}\)`)
	if printdebug {
		fmt.Printf("Matches: %v\n", matches)
	}
	output := multiplyOverDigitsInStringArray(matches)
	fmt.Printf("Final Part 1 Output: %d\n", output)
	dos_and_donts := regexMatch(data, `don\'t\(\).*?do\(\)`)
	if printdebug {
		fmt.Printf("Do and Donts: %v\n", dos_and_donts)
	}
	// Remove each and every do and dont string from the input
	for _, substring := range dos_and_donts {
		data = removeSubstringFromString(data, substring)
		if printdebug {
			fmt.Printf("After removing %s: %s\n", substring, data)
		}
	}

	// print first 5 chars
	matches = regexMatch(data, `mul\([0-9]{1,3},[0-9]{1,3}\)`)
	if printdebug {
		fmt.Printf("Matches without Do and Dont: %v\n", matches)
	}
	output = multiplyOverDigitsInStringArray(matches)
	fmt.Printf("Final Part 2 Output: %d\n", output)

}

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	data = regexp.MustCompile(`\s+`).ReplaceAll(data, []byte{})
	return string(data)
}

func regexMatch(input string, pattern string) []string {
	return regexp.MustCompile(pattern).FindAllString(input, -1)
}

func extractDigitsAndMultiply(input string) int {
	// regex to extract digits
	re := regexp.MustCompile(`\d+`)
	// There should only be two sets of digits between 0 and 999
	digits := re.FindAllString(input, -1)
	// Multiply the digits
	result := 1
	for _, digit := range digits {
		num, err := strconv.Atoi(digit)
		if err != nil {
			panic(err)
		}
		result *= num
	}
	return result
}

func multiplyOverDigitsInStringArray(input []string) int {
	result := 0
	for _, s := range input {
		result += extractDigitsAndMultiply(s)
	}
	return result
}

func removeSubstringFromString(input string, substring string) string {
	escapedSubstring := regexp.QuoteMeta(substring)
	return regexp.MustCompile(escapedSubstring).ReplaceAllString(input, "")
}
