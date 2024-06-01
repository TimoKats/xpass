package lib

import (
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
	    Info.Println(content)
	    return err
	  }
	  return errors.New("No key submitted for this locker.")
  }
  return errors.New("No locker name submitted. xpass cat-locker <<name>>")
}

