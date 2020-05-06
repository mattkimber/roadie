package dto

import (
	"roadie"
	"strings"
)

type SpritesDTO struct {
	Table                string `json:"table"`
	TemplateDirectory    string `json:"template_directory"`
	AdditionalTextField  string `json:"additional_text_field"`
	AdditionalTextFormat string `json:"additional_text_format"`
}

func (d *SpritesDTO) GetSprites() (c roadie.Sprites) {
	if len(d.Table) > 0 {
		c.Table = d.Table
	} else {
		c.Table = "table.csv"
	}
	c.TemplateDirectory = d.TemplateDirectory

	c.AdditionalTextField = d.AdditionalTextField
	if len(d.AdditionalTextField) > 0 && !strings.Contains(d.AdditionalTextFormat, "%s") {
		c.AdditionalTextFormat = d.AdditionalTextFormat + "%s"
	} else {
		c.AdditionalTextFormat = d.AdditionalTextFormat
	}

	return
}
