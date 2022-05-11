package roadie

import (
	"github.com/mattkimber/roadie/src/assets"
	"github.com/mattkimber/roadie/src/templatefunction"
	"io"
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
	t, err := assets.GetExternalTemplate([]string{s.Template}, templatefunction.Map())

	if err != nil {
		return
	}

	err = t.Execute(writer, s.Data)
	return
}
