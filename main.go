package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)
var ErrExitCLI = errors.New("exit CLI")
type cliCommand struct {
	name string
	description string
	callback func() error
}

func getInventoryCommands() map [string] cliCommand {
	return map [string] cliCommand {
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit":{
			name: "exit",
			description: "Exits the application",
			callback: commandExit,
		},
	}
}


func commandHelp() error {
	command := getInventoryCommands()
	fmt.Println("\n\nWelcome to the Pokedex!")
	fmt.Print("Usage: \n\n")

	for _, value := range command {
		fmt.Println(value.name, ":", value.description)
	}
	fmt.Print("\n\n")
	return nil
}

func commandExit() error {
	return ErrExitCLI
}


func main(){

	commands := getInventoryCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		commandName := scanner.Text()
		if command, ok := commands[commandName]; ok {
			err := command.callback()
			if err == ErrExitCLI {
				break
			}
		} 
	}
}