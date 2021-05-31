package cmd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSortExtensions(t *testing.T) {
	extCount := map[string]int{
		"md":  1,
		"go":  3,
		"txt": 2,
	}
	sortedExtensions := []kv{
		{Key: "go", Value: 3},
		{Key: "txt", Value: 2},
		{Key: "md", Value: 1},
	}

	assert.True(t, cmp.Equal(sortedExtensions, sortExtensions(extCount)))
}
