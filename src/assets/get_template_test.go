package assets

import (
	"testing"
)

func TestGetInput(t *testing.T) {
	input, err := GetInput("testdata/input.txt")
	if err != nil {
		t.Errorf("error getting input: %v", err)
	}

	if input != "Hello world" {
		t.Errorf("input not equal to expected")
	}
}
