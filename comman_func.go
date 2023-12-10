package main

import (
	"fmt"
	"strconv"
)


func clearTerminal(){
	CallClear()
}


func CheckSelectedOption(option int) bool {
	_, exists := options[option]
	return exists
}


func GetTotalTasks(tasks []Task) int{
	return len(tasks)
}


func strToInt(option string) int{
	i, err := strconv.Atoi(option)
    if err != nil {
        return 0
    }
	return i
}


func GoBack(){
	fmt.Println("Go Back")
}

func GoToMainMenu(){
	fmt.Println("Go To Main Menu")
}

func waitForInput(msg string){
	fmt.Printf("\n"+msg)
	fmt.Scanln()
}