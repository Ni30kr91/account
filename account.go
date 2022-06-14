package main

import "fmt"

func main() {
	input := [][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
		{2, 7, 9},
	}
	fmt.Println(input)
	var sum int
	arr := []int{}
	for i := 0; i < len(input); i++ {
		sum = 0
		for j := 0; j < len(input[i]); j++ {

			sum = sum + input[i][j]
		}
		fmt.Println(sum)
		arr = append(arr, sum)
	}
	fmt.Println(arr)
	largest := arr[0]
	for k := 0; k < len(arr); k++ {
		if largest < arr[k] {
			largest = arr[k]
		}
	}
	fmt.Println("largest wealth is:", largest)
}
