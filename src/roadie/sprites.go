package roadie

import (
	"assets"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"templatefunction"
	"text/template"
)

type TemplateMap map[string]*template.Template

type Sprites struct {
	Table                string
	TemplateDirectory    string
	EncounteredStrings   []LanguageString
	AdditionalTextField  string
	AdditionalTextFormat string
	NestableTemplates    []string
	Globals              map[string]string
}

func (s *Sprites) Write(w io.Writer) (err error) {
	data, err := getDataFromCsvFile(s)
	if err != nil {
		return
	}

	fields, templates, err := getFields(data, s.AdditionalTextField)
	if err != nil {
		return
	}

	s.EncounteredStrings = make([]LanguageString, 0, len(data)-1)
	for _, d := range data[1:] {
		if err = processDataLine(w, d, fields, templates, s); err != nil {
			return
		}
	}

	return
}

func processDataLine(w io.Writer, dataLine []string, fields []string, templates TemplateMap, s *Sprites) (err error) {
	templateData := make(map[string]interface{})

	for i, f := range dataLine {
		templateData[fields[i]] = f
	}

	templateData["globals"] = s.Globals

	str := LanguageString{Name: "STR_NAME_" + strings.ToUpper(templateData["id"].(string)), Value: templateData["name"].(string)}
	s.EncounteredStrings = append(s.EncounteredStrings, str)

	templateData["name_string"] = "string(" + str.Name + ")"
	if len(s.AdditionalTextField) > 0 && len(templateData[s.AdditionalTextField].(string)) > 0 {
		additional := LanguageString{
			Name:  "STR_NAME_" + strings.ToUpper(s.AdditionalTextField) + "_" + strings.ToUpper(templateData["id"].(string)),
			Value: fmt.Sprintf(s.AdditionalTextFormat, templateData[s.AdditionalTextField]),
		}
		s.EncounteredStrings = append(s.EncounteredStrings, additional)

		templateData["additional_text_string"] = "string(" + additional.Name + ")"
	}

	templateName := templateData["template"].(string)
	templateFile := s.TemplateDirectory + "/" + templateName + ".tmpl"

	if err = s.ensureTemplate(templates, templateName, templateFile); err != nil {
		return
	}

	if err = templates[templateName].Execute(w, templateData); err != nil {
		return
	}
	return
}

func (s *Sprites) ensureTemplate(templates TemplateMap, templateName string, filename string) error {
	if _, ok := templates[templateName]; !ok {
		filenames := append([]string{filename}, s.NestableTemplates...)
		tpl, err := assets.GetExternalTemplate(filenames, templatefunction.Map())
		if err != nil {
			return err
		}

		templates[templateName] = tpl
	}
	return nil
}

func getFields(data [][]string, textField string) (fields []string, templates TemplateMap, err error) {
	fields = make([]string, len(data[0]))
	templates = make(TemplateMap)

	var templateFound, idFound, nameFound, textFieldFound bool

	for i, f := range data[0] {
		// CSVs found in the wild may have BOM in the header line
		fields[i] = strings.Trim(f, " \xEF\xBB\xBF")
		if fields[i] == "template" {
			templateFound = true
		}

		if fields[i] == "id" {
			idFound = true
		}

		if fields[i] == "name" {
			nameFound = true
		}

		if fields[i] == textField {
			textFieldFound = true
		}
	}

	if !templateFound || !idFound || !nameFound {
		log.Printf("CSV headers: %v", fields)
		err = fmt.Errorf("did not find template, name and id columns in csv file")
		return
	}

	if len(textField) > 0 && !textFieldFound {
		err = fmt.Errorf("additional text field %s specified but not found in csv file", textField)
	}

	return
}

func getDataFromCsvFile(s *Sprites) (data [][]string, err error) {
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
