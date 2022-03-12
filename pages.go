package main

import "fmt"

func newPage(name string) {
	filename := getHomeDir() + CONFIG_DOCS_PATH + name + ".md"
	newFile(filename, []byte(genFrontmatter(name)))
	// open the file with text editor after creating it
	editFile(filename)
}

func openPage(name string) {
	fmt.Println("heh")
	filename := getHomeDir() + CONFIG_DOCS_PATH + name + ".md"
	if !isFileExist(filename) {
		fmt.Println("File not found!")
		return
	}
	editFile(filename)
}
