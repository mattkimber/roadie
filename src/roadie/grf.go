package roadie

import (
	"bytes"
	"github.com/mattkimber/roadie/src/assets"
	"io"
)

type Grf struct {
	Identifier                    string
	Version, MinCompatibleVersion int
	Parameters                    Parameters
	ParameterData                 string
	Language                      string
	Name, Description             string
}

func (g Grf) Write(writer io.Writer) (err error) {
	t, err := assets.GetInternalTemplate("grf", assets.GetGrfTMPL())

	if err != nil {
		return
	}

	var buf bytes.Buffer
	if err = g.Parameters.Write(&buf); err != nil {
		return
	}

	g.ParameterData = buf.String()

	err = t.Execute(writer, g)
	return
}
