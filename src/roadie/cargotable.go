package roadie

import (
	"assets"
	"io"
)

type CargoTable struct {
	Cargo []string
}

func (c CargoTable) Write(writer io.Writer) (err error) {
	if len(c.Cargo) == 0 {
		return
	}

	t, err := assets.GetInternalTemplate("cargoTable", assets.GetCargotableTMPL())

	if err != nil {
		return
	}

	err = t.Execute(writer, c)
	return
}
