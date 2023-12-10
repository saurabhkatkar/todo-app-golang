package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)


func ShowTaskDetails(r  *bufio.Reader, task Task) (int , error){
	clearTerminal()

	fmt.Println("\nBelow is Task Detail : ")

	fmt.Printf(" Title : %s \n Description : %s \n CompletionStatus : %v \n", task.Title, task.Description, task.Completed)

	fmt.Printf("\nMore Options Below : \n 1. Mark As Completed \n 2. Update Task \n 3. Go back \n 4. Main Menu \nPlease select from above options : ")

	option, err2 := r.ReadString('\n')

	if errors.Is(err2, fs.ErrExist){
		fmt.Println(err2)
		option = ""
	}

	option = strings.TrimSpace(option)

	optionInt := strToInt(option)

	if optionInt==0 {
		return optionInt, errors.New("WRONG OPTION SELECTED")
	}

	return optionInt, nil
	
}