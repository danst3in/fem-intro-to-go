package utils

import "testing"

func TestAddition(t *testing.T) {
	expected := 8
	actual := Add(4, 4)

	if actual != expected {
		t.Errorf("The Add function does not perform addition properly: Expected - %d,  Actual - %d ", expected, actual)
	}

}
