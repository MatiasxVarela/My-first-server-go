package models

type Users struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Task struct {
	TaskName string `json:"taskName"`
}