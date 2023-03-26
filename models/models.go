package models

type Users struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Task struct {
	UserId   string `json:"userId"`
	TaskName string `json:"taskName"`
}
