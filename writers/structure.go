package writers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Prashant-sharma3012/api-generator/models"
)

// responsible for creating folder structure befotre writers start writing to them

// for now hardcoding thinsg

type Folders struct {
	Name     string   `json:"name"`
	Contents []string `json:"contents"`
}

type FolderStructure struct {
	RootFolders []Folders `json:"rootFolders"`
	RootFiles   []string  `json:"rootFiles"`
	ProjectName models.Project
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

	// create empty root files

	// create empty folders

}

func WriteStaticTemplates(templateMap map[string]string, basePath string) {

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

func ParseFileToFolderMap() map[string]string {
	var pathToTemplateMap map[string]string

	content, err := ioutil.ReadFile("constants/folderStructure.json")
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
