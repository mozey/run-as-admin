package main

import (
	"fmt"
	"os/exec"
	"crypto/md5"
	"time"
	"os"
	"path"
)

func getWinDir() string {
	out, err := exec.Command("echo %WINDIR%").Output()
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", out))
		fmt.Println(err)
	}
	return string(out)
}

func isAdmin() bool {
	// File name is a hash of the current datetime string
	t := time.Now()
	data := []byte(t.Format(time.RFC3339))
	file := fmt.Sprintf("%s.txt",
		fmt.Sprintf("%x", md5.Sum(data)))
	filePath := path.Join(getWinDir(), file)

	// If we can open a file for read / write in the system folder
	// then assume that the app is running as administrator
	f, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	defer f.Close()
	if err != nil {
		return false
	}
	return true
}

func main() {
	if !isAdmin() {
		fmt.Println("Not running as administrator")
	}
	fmt.Println("Running as administrator")
	fmt.Println("")
	fmt.Println("Press any key to continue")
    fmt.Scanln()
}

