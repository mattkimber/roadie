package dto

import (
	"bufio"
	"encoding/json"
	"github.com/mattkimber/roadie/src/roadie"
	"io/ioutil"
	"os"
)

type SetDTO struct {
	GrfDTO                        GrfDTO                        `json:"grf"`
	CargoTableDTO                 CargoTableDTO                 `json:"cargo_table"`
	TemplatesCollectionDTO        SpriteTemplatesCollectionDTO  `json:"sprite_templates"`
	DynamicTemplatesCollectionDTO DynamicTemplatesCollectionDTO `json:"dynamic_templates"`
	SpritesDTO                    SpritesDTO                    `json:"sprites"`
	StaticTemplatesDTO            StaticTemplatesDTO            `json:"static_templates"`
	FixedStringsDTO               FixedStringsDTO               `json:"fixed_strings"`
}

func (d *SetDTO) GetSet(version int) (s roadie.Set) {
	s.Filename = d.GrfDTO.Filename
	s.Grf = d.GrfDTO.GetGrf(version)
	s.CargoTable = d.CargoTableDTO.GetCargoTable()
	s.Templates = d.TemplatesCollectionDTO.GetSpriteTemplateList()
	s.Templates = append(s.Templates, d.DynamicTemplatesCollectionDTO.GetSpriteTemplateList()...)
	s.Sprites = d.SpritesDTO.GetSprites()
	s.StaticTemplates = d.StaticTemplatesDTO.GetStaticTemplates()
	s.FixedStrings = d.FixedStringsDTO.GetFixedStrings()
	return
}

func FromFile(filename string, version int) (s roadie.Set, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		return
	}

	var dto SetDTO
	if err = json.Unmarshal(data, &dto); err != nil {
		return
	}

	s = dto.GetSet(version)
	err = file.Close()
	return
}
