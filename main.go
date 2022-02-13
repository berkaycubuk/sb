/*
 * Copyright (c) Berkay Çubuk <berkay@berkaycubuk.com>, 2022
 */

package main

import (
	"os"
)

// Constants
const Version = "v0.2.0"

func main() {
	loadConfig()

	// Help text
	if len(os.Args) == 1 {
		commandHelp()
		return
	}

	// Command switcher
	switch os.Args[1] {
	case "new":
		commandNew()
	case "open":
		commandOpen()
	case "git":
		commandGit()
	case "help":
		commandHelp()
	}
}
