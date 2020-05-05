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

type TemplateMap map[string]*template.Template

type Sprites struct {
	Table             string
	TemplateDirectory string
}

func (s Sprites) Write(w io.Writer) (err error) {
	data, err := getDataFromCsvFile(s)
	if err != nil {
		return
	}

	fields, templates, err := getFields(data)
	if err != nil {
		return
	}

	for _, d := range data[1:] {
		if err = processDataLine(w, d, fields, templates, s.TemplateDirectory); err != nil {
			return
		}
	}

	return
}

func processDataLine(w io.Writer, dataLine []string, fields []string, templates TemplateMap, templateDir string) (err error) {
	templateData := make(map[string]string)

	for i, f := range dataLine {
		templateData[fields[i]] = f
	}

	templateData["name_string"] = "string(STR_NAME_" + strings.ToUpper(templateData["id"]) + ")"
	templateName := templateData["template"]

	if err = ensureTemplate(templates, templateName, templateDir+"/"+templateName+".tmpl"); err != nil {
		return
	}

	if err = templates[templateName].Execute(w, templateData); err != nil {
		return
	}
	return
}

func ensureTemplate(templates TemplateMap, templateName string, filename string) error {
	if _, ok := templates[templateName]; !ok {
		tpl, err := assets.GetExternalTemplate(templateName, filename)
		if err != nil {
			return err
		}
		templates[templateName] = tpl
	}
	return nil
}

func getFields(data [][]string) (fields []string, templates TemplateMap, err error) {
	fields = make([]string, len(data[0]))
	templates = make(TemplateMap)

	var templateFound, idFound bool

	for i, f := range data[0] {
		fields[i] = f
		if fields[i] == "template" {
			templateFound = true
		}

		if fields[i] == "id" {
			idFound = true
		}

	}

	if !templateFound || !idFound {
		err = fmt.Errorf("no template and id column in csv file")
		return
	}

	return
}

func getDataFromCsvFile(s Sprites) (data [][]string, err error) {
	csvFile, err := os.Open(s.Table)
	defer csvFile.Close()
	if err != nil {
		err = fmt.Errorf("could not open csv file %s: %v", s.Table, err)
		return
	}

	rd := csv.NewReader(bufio.NewReader(csvFile))
	data, err = rd.ReadAll()

	if err != nil {
		return
	}

	if len(data) < 2 {
		err = fmt.Errorf("no data in csv file %s", s.Table)
		return
	}

	return
}
