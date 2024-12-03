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

func TestAbsValDiffCols(t *testing.T) {
	tests := []struct {
		name string
		col1 []int32
		col2 []int32
		want []int32
	}{
		{
			name: "basic test",
			col1: []int32{5, 10, 15},
			col2: []int32{2, 8, 20},
			want: []int32{3, 2, 5},
		},
		{
			name: "negative numbers",
			col1: []int32{-5, 10, -15},
			col2: []int32{2, -8, -20},
			want: []int32{7, 18, 5},
		},
		{
			name: "same numbers",
			col1: []int32{1, 1, 1},
			col2: []int32{1, 1, 1},
			want: []int32{0, 0, 0},
		},
		{
			name: "empty slices",
			col1: []int32{},
			col2: []int32{},
			want: []int32{},
		},
		{
			name: "large numbers",
			col1: []int32{50558, 25393},
			col2: []int32{44088, 45650},
			want: []int32{6470, 20257},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AbsValDiffCols(tt.col1, tt.col2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AbsValDiffCols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumCol(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		expected int32
	}{
		{
			name:     "empty slice",
			input:    []int32{},
			expected: 0,
		},
		{
			name:     "single element",
			input:    []int32{5},
			expected: 5,
		},
		{
			name:     "multiple elements",
			input:    []int32{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "negative numbers",
			input:    []int32{-1, -2, 3, -4, 5},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumCol(tt.input)
			if result != tt.expected {
				t.Errorf("sumCol(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
