package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var CONFIG_DOCS_PATH = "/Documents/sb/" // based on the $HOME directory
var CONFIG_EDITOR = "nano"

func loadConfig() {
	configFilePath := getHomeDir() + "/.sbrc"
	if !isFileExist(configFilePath) {
		// fmt.Println("Config file not found! Create .sbrc in your home directory.")
		return
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Error while reading .sbrc: %v\n", err)
		return
	}

	parseConfig(string(data))
}

func parseConfig(content string) {
	words := strings.Fields(content)

	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "set":
			// check the length for property and value
			if len(words) < i+2+1 {
				return
			}
			configSetCommand(words[i+1], words[i+2])
		}
	}
}

func configSetCommand(property string, value string) {
	switch property {
	case "editor":
		CONFIG_EDITOR = value
	case "docspath":
		CONFIG_DOCS_PATH = value
	}
}
