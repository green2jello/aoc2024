package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const debugt1 = false
const debugt2 = false

const checkstring = "XMAS"
const checkstring2 = "MAS"

func main() {
	// filename := "days/day04/test2_input.txt"
	// filename := "days/day04/input.txt"
	filename := "days/day04/input.txt"
	data := readFileIntoRune(filename)
	fmt.Println(data[0])

	matches := iterateOverEachChar(data)
	fmt.Printf("Part 1 Answer: %d\n", matches)

	matches2 := iterateOverEachChar2(data)
	fmt.Printf("Part 2 Answer: %d\n", matches2)
}

func readFileIntoRune(path string) [][]rune {
	var result [][]rune
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}
		result = append(result, []rune(line))
	}
	return result
}

func iterateOverEachChar(input [][]rune) int {
	matches := 0
	for rindex, row := range input {
		for cindex, char := range row {
			if debugt1 {
				fmt.Printf("%d %d %s\n", rindex, cindex, string(char))
			}
			if char == rune(checkstring[0]) {
				matches += checkNextChars(input, rindex, cindex, 0, 1, checkstring[1:])
				matches += checkNextChars(input, rindex, cindex, 1, 0, checkstring[1:])
				matches += checkNextChars(input, rindex, cindex, -1, 0, checkstring[1:])
				matches += checkNextChars(input, rindex, cindex, 0, -1, checkstring[1:])
				matches += checkNextChars(input, rindex, cindex, 1, 1, checkstring[1:])
				matches += checkNextChars(input, rindex, cindex, -1, 1, checkstring[1:])
				matches += checkNextChars(input, rindex, cindex, 1, -1, checkstring[1:])
				matches += checkNextChars(input, rindex, cindex, -1, -1, checkstring[1:])
			}
		}
	}
	return matches
}

// Create a rune mask that is a 3x3 matrix and looks like
// M . M
// . A .
// S . S
var (
	mask1 = [][]rune{
		{'M', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'S'},
	}
	mask2 = [][]rune{
		{'S', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'M'},
	}
	mask3 = [][]rune{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}
	mask4 = [][]rune{
		{'S', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'M'},
	}
)

func createSubMatrix(input [][]rune, rindex int, cindex int) [][]rune {
	sub := make([][]rune, 3)
	for i := range sub {
		sub[i] = make([]rune, 3)
	}
	for index, dir := range []struct {
		x int
		y int
	}{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	} {
		r := rindex + dir.x
		c := cindex + dir.y
		if r >= 0 && r < len(input) && c >= 0 && c < len(input[0]) {
			sub[index/3][index%3] = input[r][c]
		}
	}
	return sub
}

func checkNextChars(input [][]rune, rindex int, cindex int, xdir int, ydir int, checkstring string) int {
	if debugt2 {
		fmt.Printf("checkNextChars %d %d %d %d %s\n", rindex, cindex, xdir, ydir, checkstring)
		fmt.Printf("current char: %s\n", string(input[rindex][cindex]))
	}
	if len(checkstring) == 0 {
		if debugt2 {
			fmt.Printf("checkNextChars match %d %d %d %d %s\n", rindex, cindex, xdir, ydir, checkstring)
		}
		return 1
	}
	rindex += ydir
	cindex += xdir

	if rindex < 0 || rindex >= len(input) || cindex < 0 || cindex >= len(input[0]) {
		if debugt2 {
			fmt.Printf("checkNextChars out of bounds %d %d %d %d %s\n", rindex, cindex, xdir, ydir, checkstring)
		}
		return 0
	}
	if input[rindex][cindex] == rune(checkstring[0]) {
		return checkNextChars(input, rindex, cindex, xdir, ydir, checkstring[1:])
	}
	if debugt2 {
		fmt.Printf("checkNextChars no match %d %d %d %d %s\n", rindex, cindex, xdir, ydir, checkstring)
	}
	return 0
}

func iterateOverEachChar2(input [][]rune) int {
	matches := 0
	masks := [][][]rune{mask1, mask2, mask3, mask4}

	for rindex, row := range input {
		for cindex, char := range row {
			if debugt2 {
				fmt.Printf("%d %d %s\n", rindex, cindex, string(char))
			}
			if char == rune(checkstring2[1]) {
				// Check if on the edge of the matrix, if so continue
				if rindex == 0 || rindex == len(input)-1 || cindex == 0 || cindex == len(row)-1 {
					continue
				}

				// Create a sub-matrix around the current character
				sub := createSubMatrix(input, rindex, cindex)

				// Try all mask orientations
				for _, mask := range masks {
					subMatches := applyRuneMask(sub, mask)

					if subMatches > 0 {
						matches++
						break // Stop checking other masks once a match is found
					}
				}
			}
		}
	}
	return matches
}

// func prettyPrintRuneGrid(grid [][]rune) string {
// 	var sb strings.Builder
// 	for _, row := range grid {
// 		sb.WriteString(string(row))
// 		sb.WriteString("\n")
// 	}
// 	return sb.String()
// }

func applyRuneMask(matrix [][]rune, mask [][]rune) int {
	matches := 0
	// Ensure the matrix and mask have the same dimensions
	if len(matrix) != len(mask) || len(matrix[0]) != len(mask[0]) {
		return 0
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			// Skip the center of the mask (which is typically used as a reference point)
			if r == len(matrix)/2 && c == len(matrix[0])/2 {
				continue
			}

			// If mask has a non-dot rune, check for match
			if mask[r][c] != '.' {
				// Ensure the mask character matches the matrix character
				if mask[r][c] != matrix[r][c] {
					return 0 // If any non-dot mask character doesn't match, return 0
				}
				matches++
			}
		}
	}
	return matches
}
