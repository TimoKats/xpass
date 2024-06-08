package main

import (
	xps "xpass/lib"
	"errors"
	"strings"
)

func exitCommand(command string) bool {
	filteredCommand := strings.ReplaceAll(command, " ", "")
	if filteredCommand == "q\n" {
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
		case "init":
			return xps.InitLocker(command)
		case "cat":
			return xps.CatLocker(command)
		case "info":
			xps.Infobox()
		default:
			return errors.New("Command not supported")
	}
	return nil 
}

func main() {
	command := ""
	for true {
		command = xps.InsecureInput("Please enter your command: ")
		if len(xps.LockerPath) == 0 {
			xps.Error.Println("XPASS_LOCKER env variable is not set.")
			break
		}
		if exitCommand(command) {
			xps.Info.Println("Exiting xpass...")
			break
		}
		if len(command) > 1 {
			parseErr := parseCommand(strings.Fields(command))
			if parseErr != nil {
				xps.Error.Println(parseErr)
			}
		}
	}
}

