package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)

		}

		command = strings.TrimSpace(command[:len(command)-1])
		inbuiltCommands := []string{"exit", "echo", "type"}

		if strings.TrimSpace(command) == "exit" {
			os.Exit(0)
		}

		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		}

		if strings.HasPrefix(command, "type ") {
			if slices.Contains(inbuiltCommands, command[5:]) {
				fmt.Printf("%s is a shell builtin\n", command[5:])
				continue
			}

			// search for command in PATH
			fileName := strings.TrimSpace(command[5:])

			path, err := exec.LookPath(fileName)
			if err != nil {
				fmt.Printf("%s: not found\n", fileName)
				continue
			}

			fmt.Printf("%s is %s\n", fileName, path)
			continue
		}

		if command != "" {
			fmt.Printf("%s: command not found\n", command)
		}

	}
}
