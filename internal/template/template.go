package template

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

func (t Template) LoadValuesFile() (map[string]interface{}, error) {
	var values map[string]interface{}
	data, err := ioutil.ReadFile(t.Values)
	if err != nil {
		return values, err
	}

	if bytes.HasPrefix(data, []byte("{")) {
		err = json.Unmarshal(data, &values)
	} else {
		err = yaml.Unmarshal(data, &values)
	}

	if err != nil {
		return values, err
	}

	return values, nil
}

func (t TemplateEntry) LoadTemplateFile() (string, error) {
	data, err := ioutil.ReadFile(t.Template)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (t TemplateEntry) Execute(wr io.Writer, values map[string]interface{}) error {
	tmpltContent, err := t.LoadTemplateFile()
	if err != nil {
		return err
	}

	parsed, err := template.New("tmplt").Parse(tmpltContent)
	if err != nil {
		return err
	}

	err = parsed.Execute(wr, values)
	if err != nil {
		return err
	}

	return nil
}

func (t TemplateEntry) WriteToFile(values map[string]interface{}) error {
	f, err := os.Create(t.Output)
	if err != nil {
		return err
	}

	defer f.Close()

	err = t.Execute(f, values)
	if err != nil {
		return err
	}

	return nil
}
