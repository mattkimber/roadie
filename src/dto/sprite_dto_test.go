package dto

import (
	"reflect"
	"roadie"
	"testing"
)

func TestSpriteDTO_GetSprites(t *testing.T) {
	expected := roadie.Sprites{
		Table:             "table.csv",
		TemplateDirectory: "output",
	}

	dto := SpritesDTO{
		TemplateDirectory: "output",
	}

	if result := dto.GetSprites(); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
