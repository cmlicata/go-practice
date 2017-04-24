package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Enter a number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", f)
	}

	isEmpty := folderIsEmpty("/Users/ChristopherLicata/justgo")
	println(isEmpty)

}

func folderIsEmpty(path string) bool {
	// What if we were given path to a file or smth?
	fi, err := os.Stat(path)
	abortIfErr(err)
	if fi.IsDir() == false {
		err := errors.New("Destination path `" + path + "` is not a folder!")
		abortIfErr(err)
	}

	f, err := os.Open(path)
	defer f.Close()

	filesToIgnore := map[string]bool{
		".gitignore": true,
		".git":       true,
		"README.md":  true,
	}

	files, err := f.Readdirnames(0)

	if len(files) > len(filesToIgnore) {
		err := errors.New("Dir`" + path + "` is not a folder!")
		abortIfErr(err)
	}

	for _, file := range files {
		_, exists := filesToIgnore[file]
		if !exists {
			err := errors.New("Destination path `" + path + "` is not a folder!")
			abortIfErr(err)
		}
	}
	filesToIgnore
	if err == io.EOF {
		return true
	}
	// If not already exited, scanning children must have errored-out
	abortIfErr(err)
	return false
}

// Simple exit if error, to avoid putting same 4 lines of code in too many places
func abortIfErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

