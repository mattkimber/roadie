package dto

import (
	"reflect"
	"roadie"
	"testing"
)

func TestStaticTemplatesDTO_GetStaticTemplates(t *testing.T) {
	expected := roadie.StaticTemplates{
		roadie.StaticTemplate{
			Template: "test",
			Data:     map[string]string{"foo": "bar"},
		},
	}

	dto := StaticTemplatesDTO{
		StaticTemplateDTO{
			Template: "test",
			Data:     map[string]string{"foo": "bar"},
		},
	}

	if result := dto.GetStaticTemplates(); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
