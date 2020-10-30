package main

import (
	"fmt"

	"github.com/Prashant-sharma3012/api-generator/writers"
)

type routeData struct {
	ControllerName string
	RouteName      string
}

func main() {
	// tpl, err := static.GetTemplate("./templates/static/indexjs.template")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// tpl.Execute(os.Stdout, nil)

	// rd := routeData{
	// 	ControllerName: "Student",
	// 	RouteName:      "student",
	// }

	// err1 := dynamic.CreateFileFromTemplate("./templates/dynamic/route.template", rd)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	projectDetails := writers.ParseSampleJson()
	folderStructure := writers.ParseFolderStructure()
	folderStructure.AddProjectDetails(projectDetails)

	fmt.Println(projectDetails)
	fmt.Println(folderStructure)

	// filesAndFolderMap := writers.ParseFileToFolderMap()

	folderStructure.CreateEmptyStructure()
}
