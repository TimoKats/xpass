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

