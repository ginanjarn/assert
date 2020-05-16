package assert_test

import (
	"testing"

	"github.com/ginanjarn/assert"
)

func TestEqual(t *testing.T) {
	tcase := []struct {
		x interface{}
		y interface{}
	}{
		{1, 1},
		{1, 2},
		{"a", "a"},
		{"a", "b"},
		{struct{ a string }{"ho"}, struct{ a string }{"ho"}},
		{struct{ a string }{"ho"}, struct{ a string }{"hi"}},
	}
	for _, tc := range tcase {
		assert.Equal(t, tc.x, tc.y)
	}
}
