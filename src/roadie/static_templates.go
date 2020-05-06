package roadie

import (
	"assets"
	"io"
	"templatefunction"
)

type StaticTemplate struct {
	Template string
	Data     map[string]string
}

type StaticTemplates []StaticTemplate

func (s *StaticTemplates) Write(writer io.Writer) (err error) {
	for _, st := range *s {
		if err = st.Write(writer); err != nil {
			return
		}
	}

	return
}

func (s *StaticTemplate) Write(writer io.Writer) (err error) {
	t, err := assets.GetExternalTemplate(s.Template, s.Template, templatefunction.Map())

	if err != nil {
		return
	}

	err = t.Execute(writer, s.Data)
	return
}
