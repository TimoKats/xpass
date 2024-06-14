package lib

import(
	"golang.org/x/term"
	"syscall"
	"bufio"
	"fmt"
	"os"
)

var Cyan = "\033[36m" 

func InsecureInput(message string) string {
	in := bufio.NewReader(os.Stdin)
	fmt.Print(Cyan + "ACTION:   " + Reset + message)
	command, _ := in.ReadString('\n')
	return string(command)
}

func SecureInput(message string) string {
	fmt.Print(Cyan + "ACTION:   " + Reset + message)
	bytepw, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	return string(bytepw)
}


func Infobox() {
	fmt.Println("\n---\n")
	fmt.Println(Blue + "   _  __    ____     ___    _____   _____" + Reset)
	fmt.Println(Magenta + "  | |/ /   / __ \\   /   |  / ___/  / ___/" + Reset)
	fmt.Println(Red + "  |   /   / /_/ /  / /| |  \\__ \\   \\__ \\ " + Reset)
	fmt.Println(Yellow + " /   |   / ____/  / ___ | ___/ /  ___/ / " + Reset)
	fmt.Println(Yellow + "/_/|_|  /_/      /_/  |_|/____/  /____/  " + Reset)
	fmt.Println("\n---\n")
	fmt.Println("Description : Xpass is a basic password manager written in Go.")
	fmt.Println("Author      : Timo Kats (2024)")
	fmt.Println("Version     : v0.0.1")
	fmt.Println("License     : GPLv3")
	fmt.Println("\n---\n")
}

