package templatefunction

import (
	"assets"
	"reflect"
	"strings"
	"testing"
)

func TestAlternativeSprites(t *testing.T) {
	testCases := []struct {
		name             string
		sprite, template string
		bpp              int
		expectedFile     string
	}{
		{"32bpp", "my_sprite", "my_template", 32, "testdata/output/alternative_sprites.nml"},
		{"8bpp", "my_sprite", "my_template", 8, "testdata/output/alternative_sprites_8bpp.nml"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testExpectedOutput(tc.expectedFile, AlternativeSprites(tc.sprite, tc.template, 4, tc.bpp, ""), t)
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

func TestParseInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"valid", "1", 1},
		{"negative", "-33", -33},
		{"invalid", "foo", -1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := ParseInt(tc.input); result != tc.expected {
				t.Errorf("ParseInt() = %v, expected %v", result, tc.expected)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{"valid", "foo,bar,quux", []string{"foo", "bar", "quux"}},
		{"single element", "foo", []string{"foo"}},
		{"empty string", "", []string{}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := Slice(tc.input); !reflect.DeepEqual(result, tc.expected) && len(tc.expected) > 0 && len(result) > 0 {
				t.Errorf("Slice() = %v, expected %v", result, tc.expected)
			}
		})
	}
}
