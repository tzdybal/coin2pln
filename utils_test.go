package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// normal case
func TestUniqueStrings(t *testing.T) {
	data := []string{"foo", "bar", "foobar", "bar", "foo", "foofoo", "barfoo", "foofoo"}
	sort.Strings(data)
	data = uniqueStrings(data)

	assert.Equal(t, []string{"bar", "barfoo", "foo", "foobar", "foofoo"}, data)

	// corner case - nothing to do
	assert.Equal(t, []string{"bar", "barfoo", "foo", "foobar", "foofoo"}, uniqueStrings(data))
}

// corner cases
func TestUniqueStringsCornerCases(t *testing.T) {
	data := uniqueStrings(nil)
	assert.Nil(t, data)

	data = uniqueStrings([]string{})
	assert.Nil(t, data)

	data = uniqueStrings([]string{"test"})
	assert.Equal(t, []string{"test"}, data)

	data = uniqueStrings([]string{"test1", "test2"})
	assert.Equal(t, []string{"test1", "test2"}, data)

	data = uniqueStrings([]string{"test1", "test1"})
	assert.Equal(t, []string{"test1"}, data)
}
