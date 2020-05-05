package roadie

import (
	"assets"
	"io"
)

type Grf struct {
	Identifier                    string
	Version, MinCompatibleVersion int
	Params                        string
	Language                      string
}

func (g Grf) Write(writer io.Writer) (err error) {
	t, err := assets.GetTemplate("grf", assets.GetGrfTMPL())

	if err != nil {
		return
	}

	err = t.Execute(writer, g)
	return
}
