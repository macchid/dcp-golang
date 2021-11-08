package main

import "fmt"

/**
 * Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
 * For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
 * Bonus: Can you do this in one pass?
 */
func sumOfTwo(nums []int, sum int) bool {

	needs := make(map[int]bool)

	for _, n := range nums {
		// If n is needed by any other number, then return ok
		if _, ok := needs[n]; ok {
			return true
		}

		// If n is not needed by anyone, register what number n needs to complete sum.
		needs[sum-n] = true
	}

	// We never found anything.

	return false
}

func main() {
	fmt.Println(sumOfTwo([]int{10, 15, 3, 7}, 14))
}
