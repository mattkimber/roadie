package dto

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"roadie"
)

type SetDTO struct {
	GrfDTO                 GrfDTO                       `json:"grf"`
	CargoTableDTO          CargoTableDTO                `json:"cargo_table"`
	TemplatesCollectionDTO SpriteTemplatesCollectionDTO `json:"sprite_templates"`
	SpritesDTO             SpritesDTO                   `json:"sprites"`
	StaticTemplatesDTO     StaticTemplatesDTO           `json:"static_templates"`
}

func (d *SetDTO) GetSet(version int) (s roadie.Set) {
	s.Filename = d.GrfDTO.Filename
	s.Grf = d.GrfDTO.GetGrf(version)
	s.CargoTable = d.CargoTableDTO.GetCargoTable()
	s.Templates = d.TemplatesCollectionDTO.GetSpriteTemplateList()
	s.Sprites = d.SpritesDTO.GetSprites()
	s.StaticTemplates = d.StaticTemplatesDTO.GetStaticTemplates()
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
