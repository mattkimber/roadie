package roadie

import (
	"fmt"
	"github.com/mattkimber/roadie/src/assets"
	"io"
	"text/template"
)

type Parameter struct {
	Index             int
	Id                string
	Name              string
	Description       string
	NameString        string
	DescriptionString string
	Type              string
	DefaultValue      int
	MinValue          int
	MaxValue          int
	ValueNames        []string
}

type Parameters []Parameter

func (p *Parameter) Write(writer io.Writer) (err error) {
	var t *template.Template

	if len(p.ValueNames) > 0 && (p.MinValue != 0 || (p.MaxValue-p.MinValue)+1 > len(p.ValueNames)) {
		return fmt.Errorf("value name mapping is not consistent: min value is %d (must be 0) and there are %d names for %d values", p.MinValue, len(p.ValueNames), (p.MaxValue-p.MinValue)+1)
	}

	if p.Type == "int" {
		t, err = assets.GetInternalTemplate("intParam", "int_param.tmpl")
	} else {
		t, err = assets.GetInternalTemplate("boolParam", "bool_param.tmpl")
	}

	if err != nil {
		return
	}

	err = t.Execute(writer, p)
	return
}

func (p *Parameters) Write(writer io.Writer) (err error) {
	for _, param := range *p {
		if err = param.Write(writer); err != nil {
			return err
		}
	}
	return
}
