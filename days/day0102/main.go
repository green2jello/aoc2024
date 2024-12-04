package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// readCSV reads a CSV file and returns the records as a 2D string slice
func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// extractAndSortColumns takes a slice of string records and returns two sorted int32 columns
func extractAndSortColumns(records [][]string) ([]int32, []int32, error) {
	col1 := make([]int32, len(records))
	col2 := make([]int32, len(records))

	for i, row := range records {
		if len(row) < 2 {
			return nil, nil, fmt.Errorf("row %d has insufficient columns", i)
		}
		num1, err := strconv.ParseInt(row[0], 10, 32)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing first column at row %d: %v", i, err)
		}
		num2, err := strconv.ParseInt(row[1], 10, 32)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing second column at row %d: %v", i, err)
		}
		col1[i] = int32(num1)
		col2[i] = int32(num2)
	}

	// Sort both columns
	sort.Slice(col1, func(i, j int) bool { return col1[i] < col1[j] })
	sort.Slice(col2, func(i, j int) bool { return col2[i] < col2[j] })

	return col1, col2, nil
}

// AbsValDiffCols calculates the absolute difference between corresponding elements of two int32 slices
func AbsValDiffCols(col1 []int32, col2 []int32) []int32 {
	if len(col1) == 0 || len(col2) == 0 {
		return []int32{}
	}

	diffs := make([]int32, len(col1))
	for i, v := range col1 {
		diffs[i] = int32(math.Abs(float64(v - col2[i])))
	}
	return diffs
}

// sumCol calculates the sum of all elements in an int32 slice
func sumCol(col []int32) int32 {
	sum := int32(0)
	for _, v := range col {
		sum += v
	}
	return sum
}

// numberAppearances counts how many times each value from col1 appears in col2
func numberAppearances(col1 []int32, col2 []int32) []int32 {
	// Create a map to store counts from col2
	counts := make(map[int32]int32)
	for _, v := range col2 {
		counts[v]++
	}

	// Create result array and look up counts for each col1 value
	appearances := make([]int32, len(col1))
	for i, v := range col1 {
		appearances[i] = counts[v]
	}
	return appearances
}

// multiplyTwoCols multiplies corresponding elements from two int32 slices
func multiplyTwoCols(col1 []int32, col2 []int32) []int32 {
	// Create result array and multiply each col1 value with the corresponding col2 value
	multiplied := make([]int32, len(col1))
	for i, v := range col1 {
		multiplied[i] = v * col2[i]
	}
	return multiplied
}

func main() {
	records, err := readCSV("days/day02/input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	col1, col2, err := extractAndSortColumns(records)
	if err != nil {
		fmt.Println(err)
		return
	}
	appearances := numberAppearances(col1, col2)
	simScoreArray := multiplyTwoCols(col1, appearances)
	simScore := sumCol(simScoreArray)
	fmt.Println(simScore)
}
