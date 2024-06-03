package main

import (
	xps "xpass/lib"
	"errors"
	"strings"
)

func exitCommand(command string) bool {
	filteredCommand := strings.ReplaceAll(command, " ", "")
	if filteredCommand == "\n" {
		return true
	}
	return false
}

func parseCommand(command []string) error {
	switch command[0] {
		case "ls":
			return xps.ListObjects(command)
		case "add":
			return xps.AddCredentials(command)
		case "get":
			return xps.GetCredentials(command)
		case "unlock":
			return xps.AddKey(command)
		case "create-locker":
			return xps.InitLocker(command)
		case "cat-locker":
			return xps.CatLocker(command)
		default:
			return errors.New("Command not supported")
	}
	return nil 
}

func main() {
	command := ""
	for true {
		command = xps.InsecureInput("Please enter your command: ")
		if exitCommand(command) {
			xps.Info.Println("Exiting xpass...")
			break
		}
		parseErr := parseCommand(strings.Fields(command))
		if parseErr != nil {
			xps.Error.Println(parseErr)
		}
	}
}

