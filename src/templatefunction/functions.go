package templatefunction

import (
	"assets"
	"bytes"
	"fmt"
	"text/template"
)

type AlternativeSpritesInput struct {
	Sprite   string
	Template string
	Zoom     int
}

func AlternativeSprites(sprite, template string, zoom int) string {
	t, err := assets.GetInternalTemplate("intParam", assets.GetAlternative_spritesTMPL())
	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, AlternativeSpritesInput{Sprite: sprite, Template: template, Zoom: zoom}); err != nil {
		return fmt.Sprintf("%v", err)
	}

	return buf.String()
}

func Map() template.FuncMap {
	return template.FuncMap{
		"altsprites": AlternativeSprites,
	}
}
