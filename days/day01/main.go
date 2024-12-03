package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

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
		if len(row) >= 2 {
			num1, _ := strconv.ParseInt(row[0], 10, 32)
			num2, _ := strconv.ParseInt(row[1], 10, 32)
			col1[i] = int32(num1)
			col2[i] = int32(num2)
		}
	}

	// Sort both columns
	sort.Slice(col1, func(i, j int) bool { return col1[i] < col1[j] })
	sort.Slice(col2, func(i, j int) bool { return col2[i] < col2[j] })

	return col1, col2, nil
}

// readCSV reads a CSV file and returns the records

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

func sumCol(col []int32) int32 {
	sum := int32(0)
	for _, v := range col {
		sum += v
	}
	return sum
}

func main() {
	records, err := readCSV("days/day01/input.csv")
	if err != nil {
		panic(err)
	}

	col1, col2, err := extractAndSortColumns(records)
	if err != nil {
		panic(err)
	}
	PrintHeadSlice(col1)
	PrintHeadSlice(col2)

	diffs := AbsValDiffCols(col1, col2)
	PrintHeadSlice(diffs)

	total := sumCol(diffs)
	fmt.Println(total)
}
