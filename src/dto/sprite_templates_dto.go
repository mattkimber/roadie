package dto

import (
	"math"
	"roadie"
)

type SpriteTemplatesDTO struct {
	Names          []string  `json:"names"`
	LocationScales []float64 `json:"location_scales"`
	SizeScales     []float64 `json:"size_scales"`
	OffsetScales   []float64 `json:"offset_scales"`
	Locations      [][]int   `json:"locations"`
	Sizes          [][]int   `json:"sizes"`
	Offsets        [][]int   `json:"offsets"`
}

type SpriteTemplatesCollectionDTO []SpriteTemplatesDTO

func (tl SpriteTemplatesCollectionDTO) GetSpriteTemplateList() (s roadie.SpriteTemplates) {
	for _, d := range tl {
		for i, n := range d.Names {
			st := getSpriteTemplate(n, d, i)
			s = append(s, st)
		}
	}

	return
}

func getSpriteTemplate(n string, d SpriteTemplatesDTO, i int) (st roadie.SpriteTemplate) {
	st.Name = n

	locationScale := d.LocationScales[i%len(d.LocationScales)]
	sizeScale := d.SizeScales[i%len(d.SizeScales)]
	offsetScale := d.OffsetScales[i%len(d.OffsetScales)]

	st.Lines = make([]roadie.TemplateLine, len(d.Locations))

	for j := 0; j < len(d.Locations); j++ {
		st.Lines[j].X = int(math.Round(float64(d.Locations[j%len(d.Locations)][0]) * locationScale))
		st.Lines[j].Y = int(math.Round(float64(d.Locations[j%len(d.Locations)][1]) * locationScale))
		st.Lines[j].W = int(math.Round(float64(d.Sizes[j%len(d.Sizes)][0]) * sizeScale))
		st.Lines[j].H = int(math.Round(float64(d.Sizes[j%len(d.Sizes)][1]) * sizeScale))
		st.Lines[j].OffsetX = int(math.Round(float64(d.Offsets[j%len(d.Offsets)][0]) * offsetScale))
		st.Lines[j].OffsetY = int(math.Round(float64(d.Offsets[j%len(d.Offsets)][1]) * offsetScale))
	}

	return
}
