package dto

import (
	"github.com/mattkimber/roadie/src/roadie"
	"reflect"
	"testing"
)

func TestCargoTableDTO_GetCargoTable(t *testing.T) {
	expected := roadie.CargoTable{
		Cargo: []string{"ABCD", "DEFG"},
	}

	dto := CargoTableDTO{
		Cargo: []string{"ABCD", "DEFG"},
	}

	if result := dto.GetCargoTable(); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
