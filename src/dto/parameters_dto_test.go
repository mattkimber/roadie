package dto

import (
	"github.com/mattkimber/roadie/src/roadie"
	"reflect"
	"testing"
)

func TestParametersDTO_GetParameters(t *testing.T) {
	expected := roadie.Parameters{
		roadie.Parameter{
			Index:             0,
			Id:                "test_param",
			Name:              "Test Parameter",
			Description:       "Test Description",
			NameString:        "STR_PARAM_TEST_PARAM",
			DescriptionString: "STR_PARAM_TEST_PARAM_DESCRIPTION",
			Type:              "int",
			DefaultValue:      2,
			MinValue:          1,
			MaxValue:          3,
			ValueNames:        []string{"one", "two", "three"},
		},
	}

	dto := ParametersDTO{
		ParameterDTO{
			Id:           "test_param",
			Name:         "Test Parameter",
			Description:  "Test Description",
			Type:         "int",
			DefaultValue: 2,
			MinValue:     1,
			MaxValue:     3,
			ValueNames:   []string{"one", "two", "three"},
		},
	}

	if result := dto.GetParameters(); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
