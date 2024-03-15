package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println("the positions", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	res := []int{}

	for i, num := range nums {
		for j, num2 := range nums {
			//a number cannot be repeated
			if i == j {
				continue
			}

			sum := num + num2

			if sum == target {
				res = append(res, i, j)

				break
			}
		}

		if len(res) > 0 {
			break
		}
	}

	return res
}
