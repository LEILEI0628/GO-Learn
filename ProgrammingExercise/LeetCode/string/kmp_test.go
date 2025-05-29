package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKMP(t *testing.T) {
	text := "ababcabcabababd"
	pattern := "ababd"
	index := kmpSearch(text, pattern)
	assert.Equal(t, 10, index, "Expected index should be 10")
}
