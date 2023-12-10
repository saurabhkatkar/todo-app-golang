package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

func CreateNewTask(id int, task , description string, completed bool) Task {
	return Task{ID: id, Title: task, Description: description, Completed: completed}
}


func CreateTask(r *bufio.Reader,tasks []Task) []Task {
	clearTerminal()

	id := GetTotalTasks(tasks)+1

	fmt.Println("Enter your Task Title below : ")

	task, err := r.ReadString('\n')

	if errors.Is(err, fs.ErrExist){
		fmt.Println(err)
		task = ""
	}

	task = strings.TrimSpace(task)

	if(task == ""){
		return tasks
	}


	fmt.Println("Enter your Task Description below : ")

	description, err2 := r.ReadString('\n')

	if errors.Is(err2, fs.ErrExist){
		fmt.Println(err2)
		description = ""
	}

	description = strings.TrimSpace(description)

	newTask := CreateNewTask(id, task, description, false)

	tasks = append(tasks, newTask)

	WriteToFile(tasks)

	return tasks
}
