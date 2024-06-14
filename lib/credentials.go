package lib

import (
  "golang.design/x/clipboard"
  "strings"
  "errors"
  "os"
)

func findCredentials(lockerContent string, credentialsId string) (string, error) {
  credentials := strings.Split(lockerContent, "\n")
  clipboard.Init()
  for _, credential := range credentials {
    credentialFields := strings.Split(credential, "\t")
    if credentialFields[0] == credentialsId {
      Warn.Println("username: ", credentialFields[1])
      Warn.Println("password: ", credentialFields[2])
      clipboard.Write(clipboard.FmtText, []byte(credentialFields[2]))
      Warn.Println("Password copied to clipboard.")
      return "", nil
    }
  }
  Error.Println("Credentials not found in locker.")
  return "", nil
}

func GetCredentials(arguments []string) error {
  if len(arguments) == 3 {
	  key, ok := keys[arguments[1]]
	  if ok {
      filename := LockerPath + "/" + arguments[1] + ".aes"
	    content, err := DecryptRead(filename, key)
	    findCredentials(content, arguments[2])
	    return err
	  }
	  return errors.New("No key submitted for this locker.")
  }
  return errors.New("No locker name submitted: get <<lockername>> <<id>>")
}

func appendCredential(content string, credentialId string) string {
  username := InsecureInput("Username for this id: ")
  username = strings.ReplaceAll(username, "\n", "")
  password := InsecureInput("Password for this id: ")
  password = strings.ReplaceAll(password, "\n", "")
  return content + credentialId + "\t" + username + "\t" + password + "\n"
}

func AddCredentials(arguments []string) error {
  if len(arguments) == 3 {
	  key, ok := keys[arguments[1]]
	  if ok {
      filename := LockerPath + "/" + arguments[1] + ".aes"
	    content, err := DecryptRead(filename, key)
	    newContent := appendCredential(content, arguments[2])
      os.Remove(filename) // Can't overwrite encrypted files
	    EncryptWrite(arguments[1], newContent)
	    return err
	  }
	  return errors.New("No key submitted for this locker.")
  }
  return errors.New("No locker name submitted: add <<lockername>> <<id>>")
}

