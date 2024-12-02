package main

import (
	"reflect"
	"strings"
	"testing"
)

// TestReadCSVFromBuffer tests our CSV reading functionality using string inputs instead of actual files.
// This makes our tests faster and more reliable since they don't depend on external files.
func TestReadCSVFromBuffer(testContext *testing.T) {
	// In Go, we can define a slice of structs to create a table of test cases.
	// Each test case contains:
	// - name: a descriptive name for the test
	// - input: the CSV data as a string
	// - want: what we expect the CSV parser to return
	// - wantErr: whether we expect this input to cause an error
	tests := []struct {
		name    string
		input   string
		want    [][]string
		wantErr bool
	}{
		{
			name: "simple two column csv",
			// Using backticks (`) allows us to write multi-line strings.
			// The indentation is just for readability in the code - it gets trimmed later.
			input: `
				3,4
				4,3
				2,5
			`,
			// want is a 2D slice (like a matrix) of strings.
			// Each inner slice represents one row of the CSV.
			want: [][]string{
				{"3", "4"}, // First row
				{"4", "3"}, // Second row
				{"2", "5"}, // Third row
			},
			wantErr: false, // We expect this input to parse successfully
		},
		{
			name:    "empty input",
			input:   "",           // Testing with empty input
			want:    [][]string{}, // Expect an empty slice of slices
			wantErr: false,
		},
		{
			name: "malformed csv",
			input: `
				3,4
				4,3,5    // This line has 3 columns instead of 2
				2,5
			`,
			// We don't need to specify 'want' here since we expect an error
			wantErr: true, // We expect this input to cause an error due to inconsistent columns
		},
	}

	// Run each test case in the table
	for _, testCase := range tests {
		// t.Run creates a sub-test with the given name
		testContext.Run(testCase.name, func(testContext *testing.T) {
			// STEP 1: Clean up the input
			// The CSV input above is indented for readability, but we need to remove that indentation
			// and any extra whitespace before testing.
			lines := strings.Split(testCase.input, "\n") // Split input into lines
			for i := range lines {
				lines[i] = strings.TrimSpace(lines[i]) // Remove whitespace from each line
			}
			cleanInput := strings.Join(lines, "\n") // Rejoin lines into a single string

			// STEP 2: Create a string reader
			// strings.NewReader turns our string into something that can be read like a file
			reader := strings.NewReader(cleanInput)

			// STEP 3: Try to parse the CSV
			got, err := readCSVFromReader(reader)

			// STEP 4: Check if we got the error status we expected
			if (err != nil) != testCase.wantErr {
				testContext.Errorf("readCSVFromReader() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}

			// STEP 5: If we didn't expect an error, verify the parsed data
			if !testCase.wantErr {
				// Check if we got the right number of rows
				if len(got) != len(testCase.want) {
					testContext.Errorf("readCSVFromReader() got %v rows, want %v rows", len(got), len(testCase.want))
					return
				}

				// Check each row's content
				for i := range got {
					// First check if this row has the right number of columns
					if len(got[i]) != len(testCase.want[i]) {
						testContext.Errorf("row %d: got %v columns, want %v columns", i, len(got[i]), len(testCase.want[i]))
						continue
					}
					// Then check each column's content
					for j := range got[i] {
						if got[i][j] != testCase.want[i][j] {
							testContext.Errorf("row %d, col %d: got %v, want %v", i, j, got[i][j], testCase.want[i][j])
						}
					}
				}
			}
		})
	}
}

func TestConvertToColumns(t *testing.T) {
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
				{1, 3},
				{2, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToColumns(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("convertToColumns() = %v, want %v", result, tt.expected)
			}
		})
	}
}
