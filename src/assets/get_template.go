package assets

import (
	"bufio"
	"github.com/mattkimber/roadie/src/builtin"
	"io"
	"os"
	"path"
	"text/template"
)

func GetInternalTemplate(name string, filename string) (t *template.Template, err error) {
	data, err := builtin.BuiltIns.ReadFile("templates/" + filename)
	if err != nil {
		return
	}

	return template.New(name).Parse(string(data))
}

func GetExternalTemplate(filenames []string, funcMap template.FuncMap) (t *template.Template, err error) {
	t = template.New(path.Base(filenames[0]))
	t.Funcs(funcMap)

	t, err = t.ParseFiles(filenames...)
	return
}

func GetInput(filename string) (output string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	data, err := io.ReadAll(bufio.NewReader(f))
	if err != nil {
		return
	}

	output = string(data)
	err = f.Close()

	return
}
