package roadie

import (
	"github.com/mattkimber/roadie/src/assets"
	"io"
)

type TemplateLine struct {
	X, Y             int
	W, H             int
	OffsetX, OffsetY int
}

type SpriteTemplate struct {
	Name  string
	Lines []TemplateLine
}

type SpriteTemplates []SpriteTemplate

func (s SpriteTemplates) Write(writer io.Writer) (err error) {
	t, err := assets.GetInternalTemplate("spriteTemplates", assets.GetSprite_templatesTMPL())

	if err != nil {
		return
	}

	err = t.Execute(writer, s)
	return
}
