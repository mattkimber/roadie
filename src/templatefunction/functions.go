package templatefunction

import (
	"assets"
	"bytes"
	"fmt"
	"strconv"
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

func Concat(a, b string) string {
	return a + b
}

func ParseInt(input string) int {
	if i, err := strconv.Atoi(input); err != nil {
		return -1
	} else {
		return i
	}
}

func Map() template.FuncMap {
	return template.FuncMap{
		"altsprites": AlternativeSprites,
		"concat":     Concat,
		"parseint":   ParseInt,
	}
}
