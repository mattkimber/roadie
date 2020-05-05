package roadie

import (
	"bufio"
	"io"
	"os"
)

type Set struct {
	Filename   string
	Grf        Grf
	CargoTable CargoTable
	Templates  SpriteTemplates
	Sprites    Sprites
}

func (s *Set) Write(w io.Writer) (err error) {
	for _, e := range []Entity{s.Grf, s.CargoTable, s.Templates, s.Sprites} {
		if err = e.Write(w); err != nil {
			return
		}
	}
	return
}

func (s *Set) Create() (err error) {
	f, err := os.Create(s.Filename)
	if err != nil {
		return
	}

	w := bufio.NewWriter(f)
	if err = s.Write(w); err != nil {
		w.Flush()
		f.Close()
		return
	}

	w.Flush()
	err = f.Close()
	return
}
