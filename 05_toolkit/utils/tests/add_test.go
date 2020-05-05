package packages

import (
	"fem-intro-to-go/05_toolkit/utils/packages"
	"testing"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	expected := 4
	actual := packages.Add(2, 2)

	if actual != expected {
		t.Errorf("Add function does not add up: Expected: %d, Actual: %d", expected, actual)
	}
}
