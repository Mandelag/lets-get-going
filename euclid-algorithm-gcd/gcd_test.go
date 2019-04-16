package euclid

import "testing"

func TestGcd(t *testing.T) {
	actual := Gcd(42, 18)
	expected := 6
	expectEquals(t, expected, actual)
}

func TestNotFound(t *testing.T) {
	actual := Gcd(13, 21)
	expected := 1
	expectEquals(t, expected, actual)
}

// simple bdd style assertion
func expectEquals(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Error("Expected", expected, "found", actual)
	}
}