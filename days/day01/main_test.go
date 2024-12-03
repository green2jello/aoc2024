package main

import (
	"os"
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {
	// Create a temporary test file
	testData := "1,2\n3,4\n5,6"
	tmpfile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(testData)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test the readCSV function
	got, err := readCSV(tmpfile.Name())
	if err != nil {
		t.Errorf("readCSV() error = %v", err)
		return
	}

	want := [][]string{
		{"1", "2"},
		{"3", "4"},
		{"5", "6"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("readCSV() = %v, want %v", got, want)
	}
}

func TestExtractAndSortColumns(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		wantCol1 []int32
		wantCol2 []int32
		wantErr  bool
	}{
		{
			name: "basic test",
			input: [][]string{
				{"3", "4"},
				{"1", "6"},
				{"2", "5"},
			},
			wantCol1: []int32{1, 2, 3},
			wantCol2: []int32{4, 5, 6},
			wantErr:  false,
		},
		{
			name:     "empty input",
			input:    [][]string{},
			wantCol1: []int32{},
			wantCol2: []int32{},
			wantErr:  false,
		},
		{
			name: "single row",
			input: [][]string{
				{"42", "24"},
			},
			wantCol1: []int32{42},
			wantCol2: []int32{24},
			wantErr:  false,
		},
		{
			name: "unsorted numbers",
			input: [][]string{
				{"99", "10"},
				{"1", "55"},
				{"33", "22"},
			},
			wantCol1: []int32{1, 33, 99},
			wantCol2: []int32{10, 22, 55},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			col1, col2, err := extractAndSortColumns(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("extractAndSortColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(col1, tt.wantCol1) {
				t.Errorf("extractAndSortColumns() col1 = %v, want %v", col1, tt.wantCol1)
			}

			if !reflect.DeepEqual(col2, tt.wantCol2) {
				t.Errorf("extractAndSortColumns() col2 = %v, want %v", col2, tt.wantCol2)
			}
		})
	}
}

func TestConvertToInt32Matrix(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected [][]int32
	}{
		{
			name:     "empty input",
			input:    [][]string{},
			expected: nil,
		},
		{
			name: "simple 2x2",
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			expected: [][]int32{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "single row",
			input: [][]string{
				{"10", "20", "30"},
			},
			expected: [][]int32{
				{10, 20, 30},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToInt32Matrix(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("convertToInt32Matrix() = %v, want %v", result, tt.expected)
			}
		})
	}
}
