package assert

import "testing"

// Equal compare
func Equal(t *testing.T, x, y interface{}) {
	if x != y {
		t.Errorf("not equal ==> want %v expected %v", x, y)
	}
}
