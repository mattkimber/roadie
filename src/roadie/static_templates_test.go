package roadie

import (
	"testing"
)

func TestStaticTemplates_Write(t *testing.T) {
	templates := StaticTemplates{
		StaticTemplate{
			Template: "testdata/input/static.tmpl",
			Data:     map[string]string{"foo": "bar"},
		},
	}

	testTemplate(t, "testdata/output/static_templates.nml", &templates)
}
