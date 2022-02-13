package main

import (
	"fmt"
	"os"
	"os/exec"
)

func commandHelp() {
	fmt.Println("Second Brain " + Version)
	fmt.Println("Developed by: Berkay Çubuk <berkay@berkaycubuk.com>")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("	sb [command] [parameters]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("	new			Create new thing")
	fmt.Println("		project			Create new project")
	fmt.Println("		area			Create new area")
	fmt.Println("		resource		Create new resource")
	fmt.Println("	open		Open a thing")
	fmt.Println("		project			Open project")
	fmt.Println("		resource		Open resource")
	fmt.Println("	git			Use git inside documents folder")
	fmt.Println("	help		Print help")
}

func commandGit() {
	docsFullPath := getHomeDir() + CONFIG_DOCS_PATH

	commandArgs := []string{"-C", docsFullPath}

	for i := 2; i < len(os.Args); i++ {
		commandArgs = append(commandArgs, os.Args[i])
	}

	if len(os.Args) == 2 {
		commandArgs = append(commandArgs, "status")
	}

	out, err := exec.Command("git", commandArgs...).Output()
	if err != nil {
		fmt.Printf("Unable to execute git: %v\n", err)
		return
	}

	fmt.Println(string(out[:]))
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

func commandOpen() {
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
		openProject(os.Args[3])
	case "resource":
		if len(os.Args) == 3 {
			fmt.Println("No resource name given, exiting.")
			return
		}
		openResource(os.Args[3])
	}
}
