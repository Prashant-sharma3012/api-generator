package writers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Prashant-sharma3012/api-generator/models"
	"github.com/Prashant-sharma3012/api-generator/templates/static"
)

// responsible for creating folder structure befotre writers start writing to them

// for now hardcoding thinsg

type Folders struct {
	Name     string   `json:"name"`
	Contents []string `json:"contents"`
}

type FolderStructure struct {
	RootFolders    []Folders `json:"rootFolders"`
	RootFiles      []string  `json:"rootFiles"`
	ProjectDetails *models.Project
}

type FolderToTemplate struct {
	FilePath     string `json:"filePath"`
	TemplateName string `json:"templateName"`
}

type FolderToTemplateData struct {
	Data []FolderToTemplate `json:"data"`
}

func (f *FolderStructure) CreateEmptyStructure() {
	// make folder with project name
	projectName := f.ProjectDetails.ProjectName
	destination := f.ProjectDetails.Destination
	basePath := destination + "/" + projectName + "/"
	// create empty root files

	_ = os.Mkdir(basePath, 0777)

	fmt.Println("Writing files...")
	fmt.Println("Base Folder: ", basePath)

	for _, fileName := range f.RootFiles {
		fullPath := basePath + fileName
		fmt.Println("Writing: ", fullPath)
		err := ioutil.WriteFile(fullPath, nil, 0777)
		if err != nil {
			fmt.Println("error writing root file")
			fmt.Println(err)
		}
	}

	for _, folderDetails := range f.RootFolders {
		folderPath := basePath + folderDetails.Name + "/"
		_ = os.Mkdir(folderPath, 0777)
		for _, fileName := range folderDetails.Contents {
			fullPath := folderPath + fileName
			fmt.Println("Writing: ", fullPath)
			err := ioutil.WriteFile(fullPath, nil, 0777)
			if err != nil {
				fmt.Println("error writing root folders")
				fmt.Println(err)
			}
		}
	}

	// create empty folders

}

func (f *FolderStructure) AddProjectDetails(projectDetails *models.Project) {
	f.ProjectDetails = projectDetails
}

func (f *FolderStructure) GetFilePaths() []string {
	var paths []string

	for _, fileName := range f.RootFiles {
		paths = append(paths, "/"+fileName)
	}

	for _, folderDetails := range f.RootFolders {
		folderPath := "/" + folderDetails.Name + "/"
		for _, fileName := range folderDetails.Contents {
			fullPath := folderPath + fileName
			paths = append(paths, fullPath)
		}
	}

	return paths
}

func WriteStaticTemplates(templateMap map[string]string, basePath string, paths []string) {
	for _, path := range paths {
		templateName := templateMap[path]
		fullFilePath := basePath + path
		templatePath := "./templates/static/" + templateName

		fmt.Println("Writing template: ", templatePath)
		fmt.Println("Writing template to file: ", fullFilePath)

		tpl, err := static.GetTemplate(templatePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		file, err1 := os.OpenFile(fullFilePath, os.O_RDWR|os.O_CREATE, 0755)
		if err1 != nil {
			fmt.Println(err1)
			return
		}

		err2 := tpl.Execute(file, nil)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
	}
}

func ParseFolderStructure() *FolderStructure {
	content, err := ioutil.ReadFile("constants/folderStructure.json")
	if err != nil {
		log.Fatal(err)
	}

	fs := FolderStructure{}

	err = json.Unmarshal(content, &fs)
	if err != nil {
		fmt.Println("error:", err)
	}

	return &fs
}

func ParseSampleJson() *models.Project {
	content, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err)
	}

	project := models.Project{}

	err = json.Unmarshal(content, &project)
	if err != nil {
		fmt.Println("error:", err)
	}

	return &project
}

func ParseFileToFolderMap() map[string]string {
	pathToTemplateMap := map[string]string{}

	content, err := ioutil.ReadFile("constants/fileToTemplateMap.json")
	if err != nil {
		log.Fatal(err)
	}

	fd := FolderToTemplateData{}

	err = json.Unmarshal(content, &fd)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, record := range fd.Data {
		pathToTemplateMap[record.FilePath] = record.TemplateName
	}

	return pathToTemplateMap
}
