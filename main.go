package main

import (
	"bufio"
	"fmt"
	"os"
)


func main()  {
	var selectedOption=0
	
	var tasks []Task = ReadJsonFile()

	r := bufio.NewReader(os.Stdin)

	for {
		
		selectedOption = HandleUserInputs(r)

		if !CheckSelectedOption(selectedOption){
			fmt.Println("You have selected wrong option.")
			continue
		}
		
		switch options[selectedOption]{
			case "create": tasks = CreateTask(r,tasks)
			case "view": ViewTask(r,tasks)
			default: os.Exit(0)
		}
	}

}