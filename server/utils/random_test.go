package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRandomInt(t *testing.T) {
	min := int64(1)
	max := int64(10)

	for i := 0; i < 100; i++ {
		result := RandomInt(min, max)
		assert.True(t, result >= min && result <= max, "RandomInt(%d, %d) = %d; want between %d and %d", min, max, result, min, max)
	}
}

func TestRandomString(t *testing.T) {
	length := 10
	randomStr := RandomString(length)

	assert.Equal(t, length, len(randomStr), "RandomString(%d) = %s; want length %d", length, randomStr, length)

	for _, c := range randomStr {
		assert.True(t, strings.Contains(alphabet, string(c)), "RandomString(%d) contains disallowed character '%c'", length, c)
	}
}
