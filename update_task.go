package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

func UpdateTask(r *bufio.Reader, task Task) Task{
	fmt.Println("Enter new Title for Task or Press [ENTER] : ")

	title, err2 := r.ReadString('\n')

	if errors.Is(err2, fs.ErrExist){
		fmt.Println(err2)
		title = ""
	}

	title = strings.TrimSpace(title)

	fmt.Println("Enter new Description or Press [ENTER] : ")

	description, err2 := r.ReadString('\n')

	if errors.Is(err2, fs.ErrExist){
		fmt.Println(err2)
		description = ""
	}

	description = strings.TrimSpace(description)

	if title==""{
		title = task.Title
	}

	if description==""{
		description = task.Description
	}

	newTask := CreateNewTask(task.ID, title, description, task.Completed)

	return newTask
}

func MarkAsComplete( task Task) Task{
	return CreateNewTask(task.ID, task.Title, task.Description, true)
}