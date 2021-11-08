package main

import (
	"fmt"
)

/*
 * Given an array of integers, find the first missing positive integer in linear time and constant space.
 * In other words, find the lowest positive integer that does not exist in the array.
 * The array can contain duplicates and negative numbers as well.
 */
func segregate(nums []int) int {
	j := 0
	for i := range nums {
		if nums[i] <= 0 {
			k := nums[i]
			nums[i] = nums[j]
			nums[j] = k
			j++
		}
	}

	return j
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}

	return n
}

func markExisting(nums []int) {
	for _, n := range nums {
		if abs(n) > 0 && abs(n)-1 < len(nums) {
			nums[abs(n)-1] = -1 * abs(nums[abs(n)-1])
		}
	}
}

func findMissing(nums []int) int {
	for i, n := range nums {
		if n > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func missingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	shift := segregate(nums)
	fmt.Printf("Numbers: %#v, shift:%v\n", nums, shift)
	markExisting(nums[shift:])
	fmt.Printf("Marked: %#v\n", nums[shift:])
	return findMissing(nums[shift:])
}

func main() {
	fmt.Println(missingPositive([]int{3, 4, -1, 1}))
	fmt.Println(missingPositive([]int{1, 2, 0}))
	fmt.Println(missingPositive([]int{3, 2, 2, 3, -1, 1}))
}
