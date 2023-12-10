package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

func ViewOptions(r *bufio.Reader, selectedTask Task, tasks []Task ) string {

	selectedOption := ""

	for{
		
		taskOption, err := ShowTaskDetails(r, selectedTask )

		if err != nil {
			waitForInput(err.Error())
			continue
		}

		selectedOption = viewTaskOptions[taskOption]

		if selectedOption == "update"{
			selectedTask = UpdateTask(r,selectedTask)
			tasks[selectedTask.ID-1] = selectedTask 
			WriteToFile(tasks)
			
		}else if selectedOption == "complete"{
			selectedTask = MarkAsComplete(selectedTask)
			tasks[selectedTask.ID-1] = selectedTask 
			WriteToFile(tasks)
			
		}else {
			break
		}

	}
	return selectedOption
}

func ViewTask(r *bufio.Reader, tasks []Task){
	for{
	
		clearTerminal()

		fmt.Println("\nBelow are your List of Tasks : ")

		for _, v := range tasks {
			status := "Pending"
			if v.Completed { 
				status = "Completed" 
			} 
			fmt.Printf(" %d. %s - %s \n",v.ID, v.Title, status)
		}

		fmt.Print("\nEnter Task id to get details or [ENTER] to Go Back : ")
		
		option, err2 := r.ReadString('\n')

		if errors.Is(err2, fs.ErrExist){
			fmt.Println(err2)
			option = ""
		}

		option = strings.TrimSpace(option)

		optionInt := strToInt(option)

		if optionInt==0 {
			return
		}

		selectedTask := tasks[optionInt-1]

		selectedOption := ViewOptions(r,selectedTask, tasks)

		if selectedOption=="back"{
			continue

		}else{
			break

		}

	}
	
}