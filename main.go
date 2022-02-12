/*
 * Copyright (c) Berkay Çubuk <berkay@berkaycubuk.com>, 2022
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const Version = "v0.1.0"
const DocumentsPath = "/Documents/sb/" // based on the $HOME directory

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

func newProject(name string) {
	filename := getHomeDir() + DocumentsPath + "projects/" + name + ".md"
	newFile(filename, []byte(genFrontmatter(name)))

	fmt.Println("Project created!")
}

func newArea(name string) {
	path := getHomeDir() + DocumentsPath + "areas/" + name
	newFolder(path)

	fmt.Println("Area created!")
}

func newResource(name string) {
	filename := getHomeDir() + DocumentsPath + "resources/" + name + ".md"
	newFile(filename, []byte(genFrontmatter(name)))

	fmt.Println("Resource created!")
}

func commandNew() {
	if len(os.Args) == 2 {
		fmt.Println("No type given, exiting.")
		return
	}

	switch os.Args[2] {
	case "project":
		if len(os.Args) == 3 {
			fmt.Println("No project name given, exiting.")
			return
		}
		newProject(os.Args[3])
	case "area":
		if len(os.Args) == 3 {
			fmt.Println("No area name given, exiting.")
			return
		}
		newArea(os.Args[3])
	case "resource":
		if len(os.Args) == 3 {
			fmt.Println("No resource name given, exiting.")
			return
		}
		newResource(os.Args[3])
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Second Brain " + Version)
		fmt.Println("Developed by: Berkay Çubuk <berkay@berkaycubuk.com>")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("	sb [command] [parameters]")
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("	new		Create new file with prebuilt frontmatter")
		return
	}

	switch os.Args[1] {
	case "new":
		commandNew()
	}
}
