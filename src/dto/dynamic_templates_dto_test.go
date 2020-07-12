package dto

import (
	"reflect"
	"roadie"
	"testing"
)

func TestDynamicTemplatesDTO_GetSpriteTemplates(t *testing.T) {
	expected := roadie.SpriteTemplates{
		{
			Name: "test_with_offset",
			Lines: []roadie.TemplateLine{
				{X: 12, Y: 12, W: 12, H: 24, OffsetX: 3, OffsetY: 3},
				{X: 16, Y: 16, W: 16, H: 32, OffsetX: 4, OffsetY: 4},
				{X: 20, Y: 20, W: 20, H: 40, OffsetX: 5, OffsetY: 5},
				{X: 24, Y: 24, W: 24, H: 48, OffsetX: 6, OffsetY: 6},
				{X: 28, Y: 28, W: 28, H: 56, OffsetX: 7, OffsetY: 7},
				{X: 32, Y: 32, W: 32, H: 64, OffsetX: 8, OffsetY: 8},
				{X: 4, Y: 4, W: 4, H: 8, OffsetX: 1, OffsetY: 1},
				{X: 8, Y: 8, W: 8, H: 16, OffsetX: 12, OffsetY: 7},
			},
		},
	}

	dto := DynamicTemplatesCollectionDTO{
		DymamicTemplatesDTO{
			BaseTemplate: SpriteTemplatesDTO{
				Names:          []string{"test"},
				LocationScales: []float64{4},
				SizeScales:     []float64{4},
				OffsetScales:   []float64{0.25},
				Locations:      [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}},
				Sizes:          [][]int{{1, 2}, {2, 4}, {3, 6}, {4, 8}, {5, 10}, {6, 12}, {7, 14}, {8, 16}},
				Offsets:        [][]int{{4, 4}, {8, 8}, {12, 12}, {16, 16}, {20, 20}, {24, 24}, {28, 28}, {32, 32}},
			},
			Overlays: []TemplateOverlayDTO{
				{
					Names:   []string{"test_with_offset"},
					Offsets: [][]int{{0, 0}, {40, 20}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
					Shift:   -2,
				},
			},
		},
	}

	if result := dto.GetSpriteTemplateList(); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
