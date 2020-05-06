package templatefunction

import (
	"assets"
	"strings"
	"testing"
)

func TestAlternativeSprites(t *testing.T) {
	testCases := []struct {
		name             string
		sprite, template string
		expectedFile     string
	}{
		{"simple case", "my_sprite", "my_template", "testdata/output/alternative_sprites.nml"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testExpectedOutput(tc.expectedFile, AlternativeSprites(tc.sprite, tc.template, 4), t)
		})
	}
}

func testExpectedOutput(filename string, actual string, t *testing.T) {
	expected, err := assets.GetInput(filename)
	if err != nil {
		t.Errorf("could not load test data: %v", err)
	}

	if strings.ReplaceAll(actual, "\r\n", "\n") != strings.ReplaceAll(expected, "\r\n", "\n") {
		t.Errorf("template output - expected:\n%s\n---\ngot:\n%s\n---", expected, actual)
	}
}

func TestConcat(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"simple test", "foo", "bar", "foobar"},
		{"empty a", "", "bar", "bar"},
		{"empty b", "foo", "", "foo"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := Concat(tc.a, tc.b); result != tc.expected {
				t.Errorf("Concat() = %v, expected %v", result, tc.expected)
			}
		})
	}
}