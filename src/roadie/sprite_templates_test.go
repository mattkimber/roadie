package roadie

import (
	"testing"
)

func TestSpriteTemplates_Write(t *testing.T) {
	templates := SpriteTemplates{

		{
			Name: "test_template",
			Lines: []TemplateLine{
				{X: 4, Y: 4, W: 4, H: 8, OffsetX: 1, OffsetY: 1},
				{X: 8, Y: 8, W: 8, H: 16, OffsetX: 2, OffsetY: 2},
				{X: 12, Y: 12, W: 12, H: 24, OffsetX: 3, OffsetY: 3},
				{X: 16, Y: 16, W: 16, H: 32, OffsetX: 4, OffsetY: 4},
				{X: 20, Y: 20, W: 20, H: 40, OffsetX: 5, OffsetY: 5},
				{X: 24, Y: 24, W: 24, H: 48, OffsetX: 6, OffsetY: 6},
				{X: 28, Y: 28, W: 28, H: 56, OffsetX: 7, OffsetY: 7},
				{X: 32, Y: 32, W: 32, H: 64, OffsetX: 8, OffsetY: 8},
			},
		},
	}

	testTemplate(t, "testdata/output/sprite_templates.nml", templates)
}
