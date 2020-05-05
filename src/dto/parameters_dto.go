package dto

import (
	"roadie"
	"strings"
)

type ParameterDTO struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"desc"`
	Type         string   `json:"type"`
	DefaultValue int      `json:"default_value"`
	MinValue     int      `json:"min_value"`
	MaxValue     int      `json:"max_value"`
	ValueNames   []string `json:"value_names"`
}

type ParametersDTO []ParameterDTO

func (d *ParametersDTO) GetParameters() (c roadie.Parameters) {
	for i, p := range *d {
		c = append(c, roadie.Parameter{
			Index:             i,
			Id:                p.Id,
			Name:              p.Name,
			Description:       p.Description,
			NameString:        "STR_PARAM_" + strings.ToUpper(p.Id),
			DescriptionString: "STR_PARAM_" + strings.ToUpper(p.Id) + "_DESCRIPTION",
			Type:              p.Type,
			DefaultValue:      p.DefaultValue,
			MinValue:          p.MinValue,
			MaxValue:          p.MaxValue,
			ValueNames:        p.ValueNames,
		})
	}

	return
}
