package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	ints := twoSum(nums, 9)
	assert.Equal(t, ints[0], 0)
	assert.Equal(t, ints[1], 1)

	nums = []int{2, 5, 5, 11}
	ints = twoSum(nums, 10)
	assert.Equal(t, ints[0], 1)
	assert.Equal(t, ints[1], 2)

	nums = []int{3, 2, 4}
	ints = twoSum(nums, 6)
	assert.Equal(t, ints[0], 1)
	assert.Equal(t, ints[1], 2)

	nums = []int{3, 3}
	ints = twoSum(nums, 6)
	assert.Equal(t, ints[0], 0)
	assert.Equal(t, ints[1], 1)
}
