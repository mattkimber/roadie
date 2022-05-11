package dto

import (
	"github.com/mattkimber/roadie/src/roadie"
	"reflect"
	"testing"
)

func TestFixedStringsDTO_GetFixedStrings(t *testing.T) {
	expected := []roadie.LanguageString{
		{Name: "TEST", Value: "value"},
		{Name: "TEST2", Value: "value 2"},
	}

	dto := FixedStringsDTO{
		"test":  "value",
		"TEST2": "value 2",
	}

	if result := dto.GetFixedStrings(); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
