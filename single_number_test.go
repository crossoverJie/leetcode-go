package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	number := singleNumber(nums)
	assert.Equal(t, number, 1)
}
