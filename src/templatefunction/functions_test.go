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
