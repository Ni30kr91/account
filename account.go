package main

import "fmt"

func main() {
	input := [][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
		{2, 7, 9},
	}

	fmt.Println(maxwealth(input))
}
func maxwealth(accounts [][]int) int {
	answer := 0
	for i := 0; i < len(accounts); i++ {
		element := accounts[i]
		sum := 0
		for j := 0; j < len(element); j++ {

			sum = sum + element[j]
		}
		if answer < sum {
			answer = sum
		}

	}
	return answer
}
