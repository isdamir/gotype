package gotype

import (
	"github.com/isdamir/assert"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet()
	set.Add(0)
	set.Add(1)
	set.Add(2)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 3, set.Len())
	set.Remove(2)
	assert.Equal(t, 2, set.Len())
	assert.False(t, set.Contains(2))
}
