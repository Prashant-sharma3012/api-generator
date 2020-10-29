package models

type Project struct {
	ProjectName string
	Models      []Model
	Dockerize   bool
}
