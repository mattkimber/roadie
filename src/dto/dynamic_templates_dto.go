package dto

import (
	"math"
	"roadie"
)

type TemplateOverlayDTO struct {
	Names   []string `json:"names"`
	Offsets [][]int  `json:"offsets"`
	Shift   int      `json:"shift"`
}

type DymamicTemplatesDTO struct {
	BaseTemplate SpriteTemplatesDTO   `json:"base"`
	Overlays     []TemplateOverlayDTO `json:"overlays"`
}

type DynamicTemplatesCollectionDTO []DymamicTemplatesDTO

func (tl DynamicTemplatesCollectionDTO) GetSpriteTemplateList() (s roadie.SpriteTemplates) {
	for _, d := range tl {
		for _, overlay := range d.Overlays {
			for i, n := range overlay.Names {
				st := getDynamicSpriteTemplate(n, d.BaseTemplate, overlay, i)
				s = append(s, st)
			}
		}
	}

	return
}

func getDynamicSpriteTemplate(n string, d SpriteTemplatesDTO, o TemplateOverlayDTO, i int) (st roadie.SpriteTemplate) {
	st.Name = n

	offsets := make([][]int, len(d.Offsets))
	for j := 0; j < len(offsets); j++ {
		offsets[j] = make([]int, 2)
		for k := 0; k < 2; k++ {
			if len(o.Offsets) > 0 {
				offsets[j][k] = d.Offsets[j][k] + o.Offsets[j%len(o.Offsets)][k]
			} else {
				offsets[j][k] = d.Offsets[j][k]
			}
		}
	}

	locationScale := d.LocationScales[i%len(d.LocationScales)]
	sizeScale := d.SizeScales[i%len(d.SizeScales)]
	offsetScale := d.OffsetScales[i%len(d.OffsetScales)]

	st.Lines = make([]roadie.TemplateLine, len(d.Locations))

	for k := 0; k < len(d.Locations); k++ {
		j := (k + len(d.Locations) - o.Shift) % len(d.Locations)
		st.Lines[k].X = int(math.Round(float64(d.Locations[j%len(d.Locations)][0]) * locationScale))
		st.Lines[k].Y = int(math.Round(float64(d.Locations[j%len(d.Locations)][1]) * locationScale))
		st.Lines[k].W = int(math.Round(float64(d.Sizes[j%len(d.Sizes)][0]) * sizeScale))
		st.Lines[k].H = int(math.Round(float64(d.Sizes[j%len(d.Sizes)][1]) * sizeScale))
		st.Lines[k].OffsetX = int(math.Round(float64(offsets[j%len(offsets)][0]) * offsetScale))
		st.Lines[k].OffsetY = int(math.Round(float64(offsets[j%len(offsets)][1]) * offsetScale))
	}

	return
}
