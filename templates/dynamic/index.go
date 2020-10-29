package dynamic

import (
	"os"
	"text/template"
)

func CreateFileFromTemplate(path string, data interface{}) error {
	tpl := template.Must(template.ParseFiles(path))
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		return err
	}

	return nil
}
