package roadie

import (
	"testing"
)

func TestStrings_Write(t *testing.T) {
	strings := Strings{
		Language: "english",
		Data: []LanguageString{
			{Name: "STR_EXAMPLE_1", Value: "test string"},
			{Name: "STR_EXAMPLE_2", Value: "{SILVER}test"},
		},
	}

	testTemplate(t, "testdata/output/strings.lng", strings)
}
