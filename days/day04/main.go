package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const debugt1 = false
const debugt2 = true

const checkstring = "XMAS"
const checkstring2 = "MAS"

func main() {
	filename := "days/day04/test2_input.txt"
	// filename := "days/day04/input.txt"
	// filename := "days/day04/input.txt"
	data := readFileIntoRune(filename)
	fmt.Println(data[0])

	matches := iterateOverEachChar(data)
	if debugt1 {
		fmt.Printf("Part 1 Answer: %d\n", matches)
	}

	matches2 := iterateOverEachChar2(data)
	if debugt2 {
		fmt.Printf("Part 2 Answer: %d\n", matches2)
	}
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

func iterateOverEachChar2(input [][]rune) int {
	matches := 0
	for rindex, row := range input {
		for cindex, char := range row {
			if debugt2 {
				fmt.Printf("%d %d %s\n", rindex, cindex, string(char))
			}
			if char == rune(checkstring2[1]) {
				// Check if on the edge of the matrix, if so continue, else create sub matrix
				if rindex == 0 || rindex == len(input)-1 || cindex == 0 || cindex == len(row)-1 {
					continue
				}
				sub := createSubMatrix(input, rindex, cindex)
				if debugt2 {
					fmt.Printf("Sub: %v\n", sub)
				}
				matches += iterateOverSubMatrices(sub)
			}
		}
	}
	return matches
}

func iterateOverSubMatrices(input [][]rune) int {
	// Only check the edges of the submatrices
	// Pretty print the submatrices
	if debugt2 {
		fmt.Println(prettyPrintRuneGrid(input))
	}
	sub_matches := 0
	// Check if 00 is M and 02 is S or 20 is S
	// Diagonal down right
	// For each corner of the submatrix, check if the character is M
	for rindex, row := range input {
		for cindex, char := range row {
			if rindex != 1 && cindex != 1 {
				if char == rune(checkstring2[0]) {
					sub_matches += checkNextChars(input, rindex, cindex, 0, 1, checkstring2[1:])
					sub_matches += checkNextChars(input, rindex, cindex, 1, 0, checkstring2[1:])
					sub_matches += checkNextChars(input, rindex, cindex, -1, 0, checkstring2[1:])
					sub_matches += checkNextChars(input, rindex, cindex, 0, -1, checkstring2[1:])
					sub_matches += checkNextChars(input, rindex, cindex, 1, 1, checkstring2[1:])
					sub_matches += checkNextChars(input, rindex, cindex, -1, 1, checkstring2[1:])
					sub_matches += checkNextChars(input, rindex, cindex, 1, -1, checkstring2[1:])
					sub_matches += checkNextChars(input, rindex, cindex, -1, -1, checkstring2[1:])
				}
			} else {
				continue
			}
		}
	}
	sub_matches += checkNextChars(input, 0, 0, 1, 1, checkstring2)
	sub_matches += checkNextChars(input, 0, 2, -1, 1, checkstring2)
	sub_matches += checkNextChars(input, 2, 0, 1, -1, checkstring2)
	if sub_matches > 0 {
		return 1
	}
	sub_matches = 0
	sub_matches += checkNextChars(input, 0, 2, -1, -1, checkstring2)
	sub_matches += checkNextChars(input, 2, 2, -1, 1, checkstring2)
	if sub_matches > 0 {
		return 1
	}

	sub_matches = 0
	sub_matches += checkNextChars(input, 2, 2, 1, -1, checkstring[1:])
	sub_matches += checkNextChars(input, 2, 0, -1, -1, checkstring[1:])
	if sub_matches > 0 {
		return 1
	}
	return sub_matches
}

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

func prettyPrintRuneGrid(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return sb.String()
}
