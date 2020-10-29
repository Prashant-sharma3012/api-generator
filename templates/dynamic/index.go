package dynamic

import (
	"os"
	"text/template"
)

func CreateRouteFileFromTemplate(path string, data interface{}) error {
	tpl := template.Must(template.ParseFiles(path))
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		return err
	}

	return nil
}
