package packages

import (
	"fem-intro-to-go/05_toolkit/utils/packages"
	"testing"
)

func TestMakeExcited(t *testing.T) {
	expected := "OMG SO EXCITING!"
	actual := packages.MakeExcited("omg so exciting")
	if actual != expected {
		t.Errorf("Function outcome was incorrect! Expected: %s, Actual: %s", expected, actual)
	}
}
