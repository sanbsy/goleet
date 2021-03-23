package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_integerBreak(t *testing.T) {
	assert.Equal(t, 1, integerBreak(2))
	assert.Equal(t, 2, integerBreak(3))
	assert.Equal(t, 4, integerBreak(4))
	assert.Equal(t, 36, integerBreak(10))
}
