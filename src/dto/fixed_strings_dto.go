package dto

import (
	"roadie"
	"strings"
)

type FixedStringsDTO map[string]string

func (t *FixedStringsDTO) GetFixedStrings() (s []roadie.LanguageString) {
	for k, v := range *t {
		s = append(s, roadie.LanguageString{
			Name:  strings.ToUpper(k),
			Value: v,
		})
	}

	return
}
