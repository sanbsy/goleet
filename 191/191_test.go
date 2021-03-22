package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingWeight(t *testing.T) {
	assert.Equal(t, 1, hammingWeight(uint32(1)))
	assert.Equal(t, 31, hammingWeight(0b11111111111111111111111111111101))
	assert.Equal(t, 1, hammingWeight(0b00000000000000000000000010000000))
	assert.Equal(t, 3, hammingWeight(0b00000000000000000000000000001011))
}
