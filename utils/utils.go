package utils

import "fmt"

// PrintHeadSlice prints the first N elements of a slice with pretty printing
func PrintHeadSlice(slice []string, n int) {
	fmt.Printf("Printing the first %d elements:\n", n)
	for i, v := range slice {
		if i >= n {
			break
		}
		fmt.Printf("Element %d: %v\n", i+1, v)
	}
	fmt.Println("--- End of Elements ---")
}

func PrintHeadSliceInt32(slice []int32, n int) {
	fmt.Printf("Printing the first %d elements:\n", n)
	for i, v := range slice {
		if i >= n {
			break
		}
		fmt.Printf("Element %d: %v\n", i+1, v)
	}
	fmt.Println("--- End of Elements ---")
}

func PrintHeadSliceBool(slice []bool, n int) {
	fmt.Printf("Printing the first %d elements:\n", n)
	for i, v := range slice {
		if i >= n {
			break
		}
		fmt.Printf("Element %d: %v\n", i+1, v)
	}
	fmt.Println("--- End of Elements ---")
}

// func sumArray(col []int32) int32 {
// 	sum := int32(0)
// 	for _, v := range col {
// 		sum += v
// 	}
// 	return sum
// }

func SumArrayBool(col []bool) int32 {
	sum := int32(0)
	for _, v := range col {
		if v {
			sum++
		}
	}
	return sum
}
