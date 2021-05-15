package main

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	var result, usedPool []int

	numsWork := make([]int, len(nums))
	copy(numsWork, nums)

	for i := 0; i < len(numsWork); i++ {
		for j := 0; j < len(numsWork); j++ {
			if j == i {
				continue
			}
			found_i := sort.SearchInts(usedPool, i)
			found_j := sort.SearchInts(usedPool, j)
			if (found_i == len(usedPool) && found_j == len(usedPool)) && numsWork[i]+numsWork[j] == target {
				result = append(result, i)
				result = append(result, j)
				usedPool = append(usedPool, i)
				usedPool = append(usedPool, j)
			}
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15, 3, 6, 1, 8}

	target := 9

	fmt.Println(twoSum(nums, target))
}
