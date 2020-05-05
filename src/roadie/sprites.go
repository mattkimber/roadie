package roadie

import (
	"assets"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type Sprites struct {
	Table             string
	TemplateDirectory string
}

func (s Sprites) Write(w io.Writer) (err error) {
	csvFile, err := os.Open(s.Table)
	defer csvFile.Close()
	if err != nil {
		err = fmt.Errorf("could not open csv file %s: %v", s.Table, err)
		return
	}

	rd := csv.NewReader(bufio.NewReader(csvFile))
	data, err := rd.ReadAll()

	if err != nil {
		return
	}

	if len(data) < 2 {
		err = fmt.Errorf("no data in csv file %s", s.Table)
		return
	}

	fields := make([]string, len(data[0]))
	templates := make(map[string]*template.Template)
	templateFound := false

	for i, f := range data[0] {
		fields[i] = f
		if fields[i] == "template" {
			templateFound = true
		}
	}

	if !templateFound {
		err = fmt.Errorf("no template column in csv file %s", s.Table)
		return
	}

	for _, d := range data[1:] {
		templateData := make(map[string]string)

		for i, f := range d {
			templateData[fields[i]] = f
		}

		templateData["name_string"] = "string(STR_NAME_" + strings.ToUpper(templateData["id"]) + ")"

		templateName := templateData["template"]
		if _, ok := templates[templateName]; !ok {
			tpl, err := assets.GetExternalTemplate(templateName, s.TemplateDirectory+"/"+templateName+".tmpl")
			if err != nil {
				return err
			}
			templates[templateName] = tpl
		}

		if err = templates[templateName].Execute(w, templateData); err != nil {
			return
		}
	}

	return
}
