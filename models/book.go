package models

type Book struct {
	Id     string `json:"id"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}
