package models

type Users struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Task struct {
	TaskName string `json:"taskName"`
}
