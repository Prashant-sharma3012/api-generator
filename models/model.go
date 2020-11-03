package models

type Model struct {
	Name   string  `json:"name"`
	Schema []Field `json:"schema"`
}
