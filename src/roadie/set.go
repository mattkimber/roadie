package roadie

import (
	"bufio"
	"io"
	"os"
)

type Set struct {
	Filename     string
	Grf          Grf
	CargoTable   CargoTable
	Templates    SpriteTemplates
	Sprites      Sprites
	LanguageData Strings
}

func (s *Set) Write(w io.Writer) (err error) {
	for _, e := range []Entity{s.Grf, s.CargoTable, s.Templates, &s.Sprites} {
		if err = e.Write(w); err != nil {
			return
		}
	}

	s.LanguageData.Language = s.Grf.Language
	s.LanguageData.Data = make([]LanguageString, 0, 2)
	s.LanguageData.Data = append(s.LanguageData.Data, LanguageString{Name: "STR_GRF_NAME", Value: s.Grf.Name})
	s.LanguageData.Data = append(s.LanguageData.Data, LanguageString{Name: "STR_GRF_DESCRIPTION", Value: s.Grf.Description})
	s.LanguageData.Data = append(s.LanguageData.Data, s.Sprites.EncounteredStrings...)

	return
}

func (s *Set) getTotalStrings() int {
	return len(s.Sprites.EncounteredStrings) + 2
}

func (s *Set) Create() (err error) {
	if err = writeToFile(s.Filename, s.Write); err != nil {
		return
	}

	if _, err := os.Stat("lang/"); os.IsNotExist(err) {
		if err := os.Mkdir("lang/", 0755); err != nil {
			panic(err)
		}
	}

	if err = writeToFile("lang/"+s.Grf.Language+".lng", s.LanguageData.Write); err != nil {
		return
	}

	return
}

func writeToFile(filename string, writeFunc func(w io.Writer) error) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}

	w := bufio.NewWriter(f)
	if err = writeFunc(w); err != nil {
		w.Flush()
		f.Close()
		return
	}

	w.Flush()
	err = f.Close()
	return
}
