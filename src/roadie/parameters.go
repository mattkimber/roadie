package roadie

import (
	"assets"
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

	if p.Type == "int" {
		t, err = assets.GetInternalTemplate("intParam", assets.GetInt_paramTMPL())
	} else {
		t, err = assets.GetInternalTemplate("boolParam", assets.GetBool_paramTMPL())
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
