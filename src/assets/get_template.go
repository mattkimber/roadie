package assets

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
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

func GetExternalTemplate(name string, filename string, funcMap template.FuncMap) (t *template.Template, err error) {
	f, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("could not open %s: %v", filename, err)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(bufio.NewReader(f))
	if err != nil {
		return
	}

	t = template.New(name)
	t.Funcs(funcMap)
	return t.Parse(string(data))
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
