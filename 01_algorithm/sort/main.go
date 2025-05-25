package main

import "fmt"

func sort(arr []int) []int {
	return nil
}

func main() {
	tests := [][]int{
		{5, 4, 3, 2, 1},
		{10, -1, 2, 8, 0},
		{5, 4, 1, 8, 7, 3, 2, 9},
		{5, 4, 1, 8},
	}
	for _, t := range tests {
		result := sort(t)
		fmt.Printf("result: %v  data: %v\n", result, t)

	}
	//type test struct {
	//	input  []int
	//	expect []int
	//}
	//
	//tests := []test{
	//	{
	//		input:  []int{3, 1, 2},
	//		expect: []int{1, 2, 3},
	//	},
	//	{
	//		input:  []int{5, 4, 3, 2, 1},
	//		expect: []int{1, 2, 3, 4, 5},
	//	},
	//	{
	//		input:  []int{10, -1, 2, 8, 0},
	//		expect: []int{-1, 0, 2, 8, 10},
	//	},
	//	{
	//		input:  []int{},
	//		expect: []int{},
	//	},
	//	{
	//		input:  []int{1},
	//		expect: []int{1},
	//	},
	//}
	//
	//for i, tc := range tests {
	//	result := sort(tc.input)
	//	fmt.Printf("Test %d - input: %v, got: %v, expected: %v\n", i+1, tc.input, result, tc.expect)
	//}
}
