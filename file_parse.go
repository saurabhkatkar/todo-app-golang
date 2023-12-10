package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

const fileName = "tasks.json"

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(err, os.ErrNotExist)
}

func WriteToFile(data []Task){
	//...................................
	//Writing struct type to a JSON file
	//...................................
	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(fileName, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

type JsonTasks []struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Completed   bool   `json:"Completed"`
}

func ReadJsonFile() []Task{
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	if !checkFileExists(fileName){
		_, e := os.Create(fileName)
		if e != nil {
			os.Exit(0)
		}
	}

	var tasks JsonTasks

	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		tasks = JsonTasks{}
	}

	err = json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println(err)

		tasks = JsonTasks{}
	}

	var parsedTasks = []Task{}

	for _, v := range tasks {
		parsedTasks = append(parsedTasks, CreateNewTask(v.ID,v.Title,v.Description, v.Completed))
	}

	return parsedTasks
}