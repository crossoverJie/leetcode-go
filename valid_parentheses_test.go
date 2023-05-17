package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {

	assert.Equal(t, isValid("([])"), true)
	assert.Equal(t, isValid("([]"), false)
	assert.Equal(t, isValid("(]"), false)
	assert.Equal(t, isValid("("), false)
}
