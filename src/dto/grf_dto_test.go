package dto

import (
	"github.com/mattkimber/roadie/src/roadie"
	"reflect"
	"testing"
)

func TestGrfDTO_GetGrf(t *testing.T) {
	expected := roadie.Grf{
		Identifier:           "TST\\01",
		Version:              3,
		MinCompatibleVersion: 2,
		Language:             "english",
		Name:                 "Test Name",
		Description:          "Test Description",
	}

	dto := GrfDTO{
		AuthorId:             "TST",
		GrfId:                1,
		Name:                 "Test Name",
		Description:          "Test Description",
		MinCompatibleVersion: 2,
	}

	if result := dto.GetGrf(3); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
