package assets

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

func GetInternalTemplate(name string, data []byte) (t *template.Template, err error) {
	rd, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return
	}

	tdata, err := ioutil.ReadAll(rd)
	if err != nil {
		return
	}

	return template.New(name).Parse(string(tdata))
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

	data, err := ioutil.ReadAll(bufio.NewReader(f))
	if err != nil {
		return
	}

	output = string(data)
	err = f.Close()

	return
}
