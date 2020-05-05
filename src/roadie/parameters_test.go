package roadie

import (
	"testing"
)

func TestParameters_Write(t *testing.T) {
	parameters := Parameters{
		Parameter{
			Index:             0,
			Id:                "param_a",
			Name:              "Test Parameter A",
			Description:       "Test Description A",
			NameString:        "STR_PARAM_PARAM_A",
			DescriptionString: "STR_PARAM_PARAM_A_DESCRIPTION",
			Type:              "int",
			DefaultValue:      1,
			MinValue:          0,
			MaxValue:          2,
			ValueNames:        []string{"one", "two", "three"},
		},
		Parameter{
			Index:             1,
			Id:                "param_b",
			Name:              "Test Parameter B",
			Description:       "Test Description B",
			NameString:        "STR_PARAM_PARAM_B",
			DescriptionString: "STR_PARAM_PARAM_B_DESCRIPTION",
			Type:              "bool",
			DefaultValue:      0,
		},
	}

	testTemplate(t, "testdata/output/parameters.nml", &parameters)
}
