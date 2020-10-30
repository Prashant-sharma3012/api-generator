package dynamic

import (
	"text/template"
)

func GetDynamicTemplate(path string) *template.Template {
	tpl := template.Must(template.ParseFiles(path))
	return tpl
}

func GetRouterTemplate() *template.Template {
	tpl := template.Must(template.ParseFiles("./templates/dynamic/route.template"))
	return tpl
}

func GetModelTemplate() *template.Template {
	tpl := template.Must(template.ParseFiles("./templates/dynamic/model.template"))
	return tpl
}

func GetCtrlTemplate() *template.Template {
	tpl := template.Must(template.ParseFiles("./templates/dynamic/controller.template"))
	return tpl
}

func GetRepoTemplate() *template.Template {
	tpl := template.Must(template.ParseFiles("./templates/dynamic/repository.template"))
	return tpl
}
