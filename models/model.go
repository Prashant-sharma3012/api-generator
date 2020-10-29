package models

type Model struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
}
