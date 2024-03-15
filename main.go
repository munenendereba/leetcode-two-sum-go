package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println("the positions", twoSumHashmap(nums, target))
}

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j, num2 := range nums {
			//a number cannot be repeated
			if i == j {
				continue
			}

			sum := num + num2

			if sum == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSumHashmap(nums []int, target int) []int {
	numbers := make(map[int]int)

	for i, num := range nums {
		wewant := target - num

		numgot, ok := numbers[wewant]

		if ok {
			return []int{i, numgot}
		}

		numbers[num] = i

	}

	return []int{}
}
