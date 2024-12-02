package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func readCSVFromReader(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)
	return reader.ReadAll()
}

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return readCSVFromReader(file)
}

// convertToInt32Matrix converts a matrix of strings to a matrix of int32 values
func convertToInt32Matrix(records [][]string) [][]int32 {
	if len(records) == 0 {
		return nil
	}

	result := make([][]int32, len(records))
	for i, row := range records {
		result[i] = make([]int32, len(row))
		for j, val := range row {
			num, _ := strconv.ParseInt(val, 10, 32)
			result[i][j] = int32(num)
		}
	}
	return result
}

func main() {
	records, err := readCSV("days/day01/input.csv")
	if err != nil {
		panic(err)
	}
	_ = records // TODO: Process the records
}
