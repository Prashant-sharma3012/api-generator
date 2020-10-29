package models

type Project struct {
	ProjectName string  `json:"projectName"`
	Models      []Model `json:"models"`
	Dockerize   bool    `json:"dockrize"`
	Destination string  `json:"destination"`
}
