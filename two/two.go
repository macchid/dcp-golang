package main

import "fmt"

/**
 * Given an array of integers, return a new array such that each element at index i
 * of the new array is the product of all the numbers in the original array except the one at i.
 */

func productWithouI(nums []int) []int {

	prod := 1
	for _, n := range nums {
		prod *= n
	}

	res := make([]int, len(nums))
	for i, n := range nums {
		res[i] = prod / n
	}

	return res
}

func main() {
	fmt.Println(productWithouI([]int{10, 9, 8, 7}))
}
