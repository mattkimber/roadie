package templatefunction

import (
	"bytes"
	"fmt"
	"github.com/mattkimber/roadie/src/assets"
	"strconv"
	"strings"
	"text/template"
)

type AlternativeSpritesInput struct {
	Sprite     string
	SpriteName string
	Template   string
	Zoom       int
	Bpp        int
}

// Slightly hacky as we originally only had the 32bpp method
func AlternativeSprites32(sprite, template string, zoom int) string {
	return AlternativeSprites(sprite, template, zoom, 32, "")
}

func AlternativeSprites8(sprite, template string, zoom int) string {
	return AlternativeSprites(sprite, template, zoom, 8, "")
}

func AlternativeSprites8N(sprite, template string, zoom int, suffix string) string {
	return AlternativeSprites(sprite, template, zoom, 8, suffix)
}

func AlternativeSprites32N(sprite, template string, zoom int, suffix string) string {
	return AlternativeSprites(sprite, template, zoom, 32, suffix)
}

func AlternativeSprites(sprite, template string, zoom int, bpp int, suffix string) string {
	t, err := assets.GetInternalTemplate("intParam", "alternative_sprites.tmpl")
	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, AlternativeSpritesInput{Sprite: sprite, Template: template, Zoom: zoom, Bpp: bpp, SpriteName: sprite + suffix}); err != nil {
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

func ToUpper(input string) string {
	return strings.ToUpper(input)
}

func Slice(input string) []string {
	return strings.Split(input, ",")
}

func Dict(values ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		key := values[i].(string)
		m[key] = values[i+1]
	}
	return m
}

func Map() template.FuncMap {
	return template.FuncMap{
		"altsprites":   AlternativeSprites32,
		"altsprites8":  AlternativeSprites8,
		"altspritesn":  AlternativeSprites32N,
		"altsprites8n": AlternativeSprites8N,
		"concat":       Concat,
		"parseint":     ParseInt,
		"slice":        Slice,
		"toupper":      ToUpper,
		"dict":         Dict,
	}
}
