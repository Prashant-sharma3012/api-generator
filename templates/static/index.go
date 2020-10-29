package static

import "text/template"

func GetTemplate(path string) (*template.Template, error) {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	return tpl, nil
}
