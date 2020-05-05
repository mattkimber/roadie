package roadie

import (
	"fmt"
	"io"
)

type LanguageString struct {
	Name  string
	Value string
}

type Strings struct {
	Language string
	Data     []LanguageString
}

func (s Strings) Write(writer io.Writer) (err error) {
	if len(s.Data) == 0 {
		return
	}

	languageID := GetLanguageId(s.Language)
	fmt.Fprintf(writer, "##grflangid 0x%02X\n", languageID)

	for _, str := range s.Data {
		fmt.Fprintf(writer, "%s :%s\n", str.Name, str.Value)
	}

	return
}
