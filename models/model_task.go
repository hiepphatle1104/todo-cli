package models

type Task struct {
	Name   string `json:"name"`
	IsDone bool   `json:"is-done"`
	TaskId int    `json:"task-id"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}
