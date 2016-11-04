package main

import (
	"fmt"
	"crypto/md5"
	"time"
	"os"
)

const winDir = "c:\\Windows"

// isAdmin tries to create a new file in the windows dir, if the file is created
// then assume the program is running as administrator and return true
func isAdmin() bool {
	 //File name is a hash of the current datetime string
	t := time.Now()
	data := []byte(t.Format(time.RFC3339))
	filePath := fmt.Sprintf("%s\\%s.txt", winDir,
		fmt.Sprintf("%x", md5.Sum(data)))

	fmt.Println(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// filePath does not exist
		fmt.Println("not exist")

		// If we can open a file for read / write in the system folder
		// then assume that the app is running as administrator
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			fmt.Println("false")
			return false
		}

		f.Close()
		err = os.Remove(filePath)
		if err != nil {
			fmt.Println(err)
			return false
		}

		fmt.Println("true")
		return true

	} else {
		// File already exists, try a different filePath
		fmt.Println("exists")
		return isAdmin()
	}
}

func main() {
	if isAdmin() {
		fmt.Println("Running as administrator")
	} else {
		fmt.Println("Not running as administrator")
	}
	fmt.Println("")
	fmt.Println("Press any key to continue")
    fmt.Scanln()
}

