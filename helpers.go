package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func editFile(filePath string) {
	cmd := exec.Command(CONFIG_EDITOR, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func isFileExist(filepath string) bool {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to get $HOME: %v\n", err)
		return ""
	}

	return dirname
}

func newFile(filename string, content []byte) {
	err := ioutil.WriteFile(filename, content, 0664)
	if err != nil {
		fmt.Printf("Unable to write file: %v\n", err)
		return
	}
}

func newFolder(foldername string) {
	err := os.Mkdir(foldername, 0755)
	if err != nil {
		fmt.Printf("Unable to create folder: %v\n", err)
		return
	}
}

func genFrontmatter(title string) string {
	frontmatter := "---\n"
	frontmatter += "title: " + title + "\n"
	frontmatter += "---\n"

	return frontmatter
}
