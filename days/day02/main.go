package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"aoc2024/utils"
)

func main() {
	day21debug := false
	day22debug := true
	fmt.Printf("Day 2 Part 1 debug mode: %t\n", day21debug)
	fmt.Printf("Day 2 Part 2 debug mode: %t\n", day22debug)
	lines, err := readLines("days/day02/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Print the first 5 elements using the utility function
	if day21debug {
		utils.PrintHeadSlice(lines, 5)
	}
	intlines := convertArrayStringToArrayOfArrayOfInts(lines)
	if day21debug {
		utils.PrintHeadSliceInt32(intlines[0], 5)
	}

	safe := processAllIntLines(intlines)
	if day21debug {
		utils.PrintHeadSliceBool(safe, 5)
	}

	total := utils.SumArrayBool(safe)
	fmt.Printf("Part 1 Answer: %d\n", total)

	safe2 := processAllIntLines2(intlines)
	if day22debug {
		utils.PrintHeadSliceBool(safe2, 6)
	}

	total2 := utils.SumArrayBool(safe2)
	fmt.Printf("Part 2 Answer: %d\n", total2)

	// Lol too stupid for part 2 and ran out of time by end of evening. Going to bed instead.
	otherTestExamples := [][]int32{
		{48, 46, 47, 49, 51, 54, 56}, // safe
		{1, 1, 2, 3, 4, 5},           // safe
		{1, 2, 3, 4, 5, 5},           // safe
		{5, 1, 2, 3, 4, 5},           // safe
		{1, 4, 3, 2, 1},              // safe
		{1, 6, 7, 8, 9},              // safe
		{1, 2, 3, 4, 3},              // safe
		{9, 8, 7, 6, 7},              // safe
	}

	safe3 := processAllIntLines2(otherTestExamples)
	fmt.Println(safe3)
}

// readLines reads a whole file into memory and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func convertStringElementToArrayofInts(element string) []int32 {
	var result []int32
	for _, numStr := range strings.Split(element, " ") {
		num, err := strconv.ParseInt(numStr, 10, 32)
		if err != nil {
			log.Fatalf("Error converting string to int32: %v", err)
		}
		result = append(result, int32(num))
	}
	return result
}

func convertArrayStringToArrayOfArrayOfInts(arr []string) [][]int32 {
	var result [][]int32
	for _, element := range arr {
		result = append(result, convertStringElementToArrayofInts(element))
	}
	return result
}

func processAllIntLines(integers [][]int32) []bool {
	var result []bool
	for _, ints := range integers {
		result = append(result, processIntLines(ints))
	}
	return result
}

func processAllIntLines2(integers [][]int32) []bool {
	var result []bool
	for _, ints := range integers {
		result = append(result, processIntLines2(ints))
	}
	return result
}

func processIntLines(integers []int32) bool {
	queueint := createQueueFromInts(integers)

	if queueint.Len() < 2 {
		return true // Not enough elements to determine trend
	}

	current := queueint.Remove(queueint.Front()).(int32)
	next := queueint.Remove(queueint.Front()).(int32)
	difference := current - next

	if difference > 3 || difference < -3 || difference == 0 {
		return false
	}

	increasing := difference < 0

	for queueint.Len() > 0 {
		current = next
		next = queueint.Remove(queueint.Front()).(int32)
		difference = current - next

		if difference > 3 || difference < -3 || difference == 0 {
			return false
		}

		if (difference < 0) != increasing {
			return false
		}
	}

	return true
}

func processIntLines2(integers []int32) bool {
	queueint := createQueueFromInts(integers)
	day22debug := false

	if queueint.Len() < 3 {
		return true // Not enough elements to determine trend
	}

	current := queueint.Remove(queueint.Front()).(int32)
	next := queueint.Remove(queueint.Front()).(int32)
	last := queueint.Remove(queueint.Front()).(int32)
	one_exception := false
	increasing := (current - next) < 0
	altincreasing := (current - last) < 0
	conincreasing := (next - last) < 0

	for {
		this_exception := false
		diff1 := current - next
		diff2 := current - last
		diff3 := next - last
		if day22debug {
			fmt.Printf("Current: %d, Next: %d, Last: %d\n", current, next, last)
			fmt.Printf("Diff1: %d, Diff2: %d\n", diff1, diff2)
			fmt.Printf("One Exception: %t, Increasing: %t, Alt Increasing: %t\n", one_exception, increasing, altincreasing)
		}

		if diff1 > 3 || diff1 < -3 || diff1 == 0 {
			if one_exception {
				return false
			}
			this_exception = true
			// This diff should check if dropping next would meet criteria
			if diff2 > 3 || diff2 < -3 || diff2 == 0 {
				// This diff should check if dropping first would meet criteria
				if diff3 > 3 || diff3 < -3 || diff3 == 0 {
					return false
				}
			}
		}

		if (diff1 < 0) != increasing {
			if one_exception {
				return false
			}
			this_exception = true
			if (diff2 < 0) != altincreasing {
				if (diff3 < 0) != conincreasing {
					return false
				}
			}
		}

		if queueint.Len() == 0 {
			break
		}
		if this_exception {
			one_exception = true
		}
		current = next
		next = last
		last = queueint.Remove(queueint.Front()).(int32)
	}

	return true
}

func createQueueFromInts(integers []int32) *list.List {
	queue := list.New()

	// Enqueue each element
	for _, num := range integers {
		queue.PushBack(num)
	}

	return queue
}
