package models

var Tasks []Task = []Task{}

type Task struct {
	Priority int
	Name     string
}
