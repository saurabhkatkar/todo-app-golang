package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
)

func HandleUserInputs(r *bufio.Reader) int {
	clearTerminal()

	fmt.Printf(" \n!------ Welcome to Todo App ------!.\n 1. Create New Task \n 2. View Task \n 3. Exit \nPlease select from above options: ")


	var option int

	_, err := fmt.Scanln(&option)

	if errors.Is(err, fs.ErrExist){
		fmt.Println(err)
		option = 0
	}

	return option
}