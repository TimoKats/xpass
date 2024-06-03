package lib

import (
  "strings"
  "errors"
  "os"
)

var lockerPath = "/home/timokats/.xpass"

func createLockerText(lockername string) string {
  return "LOCKERNAME: [" + lockername + "]\n"
}

func createLockerFolder() error {
  err := os.MkdirAll(lockerPath, os.ModePerm)
  if err != nil {
    return errors.New("Error when creating the xpass folder. Please check path.")
  }
  return nil
}

func createLockerFile(lockername string) error {
  lockerText := createLockerText(lockername)
  initKey(lockername)
  writeErr := EncryptWrite(lockername, lockerText)
  return writeErr
}

func InitLocker(arguments []string) error {
  Warn.Println(arguments)
  if len(arguments) == 2 {
    folderErr := createLockerFolder()
    lockerErr := createLockerFile(arguments[1])
    if folderErr != nil {
      return folderErr 
    } else if lockerErr != nil {
      return lockerErr 
    } 
    return nil
  }
  return errors.New("No locker name submitted. xpass create-locker <<name>>")
}

func CatLocker(arguments []string) error {
  if len(arguments) == 2 {
	  key, ok := keys[arguments[1]]
	  if ok {
      filename := lockerPath + "/" + arguments[1] + ".aes"
	    content, err := DecryptRead(filename, key)
	    Info.Println("\n\n" + content)
	    return err
	  }
	  return errors.New("No key submitted for this locker.")
  }
  return errors.New("No locker name submitted. xpass cat-locker <<name>>")
}

func listObjectsInLocker(lockername string) error {
	key, ok := keys[lockername]
	if ok {
    filename := lockerPath + "/" + lockername + ".aes"
	  lockerContent, err := DecryptRead(filename, key)
    credentials := strings.Split(lockerContent, "\n")
    for _, credential := range credentials[1:] {
      credentialFields := strings.Split(credential, "\t")
      Info.Println(credentialFields[0])
	  }
	  return err
	}
	return errors.New("No key submitted for this locker.")
}

func ListObjects(arguments []string) error {
  if len(arguments) == 1 {
    lockers, _ := os.ReadDir(lockerPath)
    for _, locker := range lockers {
      Info.Println(locker.Name()[:len(locker.Name()) - 4])
    }
    return nil
  } else if len(arguments) == 2 {
    return listObjectsInLocker(arguments[1]) 
  }
  return errors.New("Unkown command. xpass ls <<name: optional>>")
}
