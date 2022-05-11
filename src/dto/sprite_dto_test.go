package dto

import (
	"github.com/mattkimber/roadie/src/roadie"
	"reflect"
	"testing"
)

func TestSpriteDTO_GetSprites(t *testing.T) {
	expected := roadie.Sprites{
		Table:                "table.csv",
		TemplateDirectory:    "output",
		AdditionalTextField:  "text",
		AdditionalTextFormat: "%s",
		Globals:              map[string]string{"foo": "bar"},
	}

	dto := SpritesDTO{
		TemplateDirectory:   "output",
		AdditionalTextField: "text",
		Globals:             map[string]string{"foo": "bar"},
	}

	if result := dto.GetSprites(); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
