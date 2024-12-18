package main

import "fmt"

func findMax(arr []int) int {
	//Initialize max with the first element of the array
	max := arr[0]

	//Loop through the array starting from the second element
	for _, num := range arr[1:] {
		if num > max {
			max = num
		}
	}

	// return the maximum value found
	return max
}

func main() {
	arr1 := []int{1,2,3,4,5}
	arr2 := []int{-1,-2,-3,-4,-5}

	fmt.Println("Maximum in arr1:", findMax(arr1))
	fmt.Println("Maximum in arr2:", findMax(arr2))
}