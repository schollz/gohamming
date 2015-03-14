package hamming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidInput(t *testing.T) {
	assert.Equal(t, -1, Calc("aaa", "aa"), "should be -1")
	assert.Equal(t, -1, Calc("aa", "fffaa"), "should -1")
	assert.Equal(t, -1, Calc("aa", ""), "should -1")
}
func TestCalculation(t *testing.T) {
	assert.Equal(t, 0, Calc("abc", "abc"), "should be 0")
	assert.Equal(t, 1, Calc("abc", "abv"), "should be 1")
	assert.Equal(t, 3, Calc("karolin", "kathrin"), "should be 3")
	assert.Equal(t, 3, Calc("karolin", "kerstin"), "should be 3")
	assert.Equal(t, 2, Calc("1011101", "1001001"), "should be 3")
}
