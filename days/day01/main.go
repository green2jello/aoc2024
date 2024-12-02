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

// convertToColumns transforms rows of strings into columns of int32
func convertToColumns(records [][]string) [][]int32 {
	if len(records) == 0 {
		return nil
	}

	numCols := len(records[0])
	cols := make([][]int32, numCols)

	for i := 0; i < numCols; i++ {
		cols[i] = make([]int32, len(records))
	}

	for i, row := range records {
		for j, val := range row {
			num, _ := strconv.ParseInt(val, 10, 32)
			cols[j][i] = int32(num)
		}
	}
	return cols
}

func main() {
	records, err := readCSV("days/day01/input.csv")
	if err != nil {
		panic(err)
	}
	_ = records // TODO: Process the records
}
