package task

import "fmt"

// Function that runs when the sb run
func Init() {

}

func CommandListTasks() {

}

// Function to register commands
func RegisterCommands(args []string) {
	if args[1] != "task" {
		return
	}

	fmt.Println("asdfasf")
}
