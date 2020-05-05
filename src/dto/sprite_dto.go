package dto

import "roadie"

type SpritesDTO struct {
	Table             string `json:"table"`
	TemplateDirectory string `json:"template_directory"`
}

func (d *SpritesDTO) GetSprites() (c roadie.Sprites) {
	if len(d.Table) > 0 {
		c.Table = d.Table
	} else {
		c.Table = "table.csv"
	}
	c.TemplateDirectory = d.TemplateDirectory
	return
}
