package models

var Tasks []Task = []Task{}

type Task struct {
	Name   string
	Method func() any
}
