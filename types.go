package main

type Task struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Completed   bool   `json:"Completed"`
}

var options = map[int]string{
	1: "create",
	2: "view",
	3: "exit"}

var viewTaskOptions = map[int]string{
	1: "complete",
	2: "update",
	3: "back",
	4: "menu"}