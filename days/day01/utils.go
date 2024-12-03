package main

import "fmt"

// PrintHeadSlice prints the first 5 elements of a slice
func PrintHeadSlice(slice []int32) {
	for _, v := range slice[:5] {
		fmt.Println(v)
	}
}
