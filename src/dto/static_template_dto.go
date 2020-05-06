package dto

import (
	"roadie"
)

type StaticTemplateDTO struct {
	Template string            `json:"template"`
	Data     map[string]string `json:"data"`
}

type StaticTemplatesDTO []StaticTemplateDTO

func (t *StaticTemplatesDTO) GetStaticTemplates() (s roadie.StaticTemplates) {
	for _, st := range *t {
		s = append(s, roadie.StaticTemplate{
			Template: st.Template,
			Data:     st.Data,
		})
	}

	return
}
