package main

import (
	"github.com/Prashant-sharma3012/api-generator/writers"
)

type routeData struct {
	ControllerName string
	RouteName      string
}

func main() {
	projectDetails := writers.ParseSampleJson()
	folderStructure := writers.ParseFolderStructure()
	folderStructure.AddProjectDetails(projectDetails)

	filesAndFolderMap := writers.ParseFileToFolderMap()
	folderStructure.CreateEmptyStructure()

	writers.WriteStaticTemplates(filesAndFolderMap, projectDetails.ProjectName, folderStructure.GetFilePaths())
	folderStructure.CreateDynamicFiles()
}
