package assets

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"text/template"
)

func GetTemplate(name string, data []byte) (t *template.Template, err error) {
	rd, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return
	}

	tdata, err := ioutil.ReadAll(rd)
	if err != nil {
		return
	}

	return template.New("grf").Parse(string(tdata))
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
