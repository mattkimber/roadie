package roadie

import (
	"testing"
)

func TestSprites_Write(t *testing.T) {
	sprites := Sprites{
		Table:             "testdata/input/table.csv",
		TemplateDirectory: "testdata/input",
	}

	testTemplate(t, "testdata/output/sprite.nml", &sprites)
}
