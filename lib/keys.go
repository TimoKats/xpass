package lib

import (
	"errors"
)

var keys = make(map[string]string)

func AddKey(arguments []string) error {
	if len(arguments) == 2 {
		key := SecureInput("Please enter your keypass: ")
		keys[arguments[1]] = key
	} else {
		return errors.New("Not enough arguments provided.")
	}
	return nil
}

func initKey(lockername string) {
	key := SecureInput("Please enter your keypass: ")
	keys[lockername] = key
	Info.Println("Created key for locker.")
}

