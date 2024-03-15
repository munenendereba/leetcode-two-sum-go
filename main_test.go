package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums   []int
	target int
	result []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		testname := fmt.Sprintf(`testing with input "%v", "%d"`, test.nums, test.target)

		t.Run(testname, func(t *testing.T) {
			result := twoSum(test.nums, test.target)

			err := fmt.Errorf(`testing twoSum("%v","%d"), expected "%v", but got "%v"`,
				test.nums, test.target, test.result, result)

			assert.ElementsMatch(t, result, test.result, err)
		})
	}
}

func TestTwoSumHashmap(t *testing.T) {
	for _, test := range tests {
		testname := fmt.Sprintf(`testing with input "%v", "%d"`, test.nums, test.target)

		t.Run(testname, func(t *testing.T) {
			result := twoSumHashmap(test.nums, test.target)

			err := fmt.Errorf(`testing twoSum("%v","%d"), expected "%v", but got "%v"`,
				test.nums, test.target, test.result, result)

			assert.ElementsMatch(t, result, test.result, err)
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{98, 3, 87}, 90)
	}
}

func BenchmarkTwoSumHashmap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumHashmap([]int{98, 3, 87}, 90)
	}
}
