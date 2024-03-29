package dto

import (
	"fmt"
	"github.com/mattkimber/roadie/src/roadie"
)

type GrfDTO struct {
	AuthorId             string        `json:"author_id"`
	GrfId                int           `json:"grf_id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	MinCompatibleVersion int           `json:"min_compat_version"`
	Filename             string        `json:"filename"`
	Language             string        `json:"language"`
	ParametersDTO        ParametersDTO `json:"parameters"`
}

func (d *GrfDTO) GetGrf(version int) (g roadie.Grf) {
	g.Identifier = fmt.Sprintf("%s\\%02X", d.AuthorId, d.GrfId)
	g.MinCompatibleVersion = d.MinCompatibleVersion
	g.Version = version
	g.Name = d.Name
	g.Description = d.Description
	g.Parameters = d.ParametersDTO.GetParameters()

	if len(d.Language) > 0 {
		g.Language = d.Language
	} else {
		g.Language = "english"
	}

	return
}
